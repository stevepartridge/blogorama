package certs

import (
	"io/ioutil"
	"os"

	"crypto/tls"

	"github.com/stevepartridge/blogorama/pkg/file"

	"github.com/stevepartridge/blogorama/pkg/certs/insecure"
)

const (
	envTLSCert   = "TLS_PUBLIC_CERT"
	envTLSKey    = "TLS_PRIVATE_KEY"
	envTLSRootCA = "TLS_ROOT_CA"
)

func GetCert() ([]byte, error) {

	if os.Getenv(envTLSCert) != "" {
		if file.Exists(os.Getenv(envTLSCert)) {
			publicCert, err := ioutil.ReadFile(os.Getenv(envTLSCert))
			if err != nil {
				return nil, err
			}
			return []byte(publicCert), nil
		}

		return []byte(os.Getenv(envTLSCert)), nil
	}

	return []byte(insecure.Cert), nil
}

func GetKey() ([]byte, error) {

	if os.Getenv(envTLSKey) != "" {
		if file.Exists(os.Getenv(envTLSKey)) {
			privateKey, err := ioutil.ReadFile(os.Getenv(envTLSKey))
			if err != nil {
				return nil, err
			}
			return []byte(privateKey), nil
		}

		return []byte(os.Getenv(envTLSKey)), nil

	}

	return []byte(insecure.Key), nil
}

func GetCertificate() (tls.Certificate, error) {

	publicCert, err := GetCert()
	if err != nil {
		return tls.Certificate{}, err
	}

	privateKey, err := GetKey()
	if err != nil {
		return tls.Certificate{}, err
	}

	return tls.X509KeyPair(publicCert, privateKey)
}

func GetRootCA() ([]byte, error) {

	if os.Getenv(envTLSRootCA) != "" {
		if file.Exists(os.Getenv(envTLSRootCA)) {
			rootCA, err := ioutil.ReadFile(os.Getenv(envTLSRootCA))
			if err != nil {
				return nil, err
			}
			if rootCA != nil {
				return rootCA, nil
			}
		}

		return []byte(os.Getenv(envTLSRootCA)), nil
	}

	return []byte(insecure.RootCA), nil
}

func GetCertAndKey() ([]byte, []byte, error) {
	cert, err := GetCert()
	if err != nil {
		return nil, nil, err
	}

	key, err := GetKey()
	if err != nil {
		return nil, nil, err
	}

	return cert, key, nil
}
