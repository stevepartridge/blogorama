package certs


import (
	"errors"
	"fmt"
)

var (
	ErrRootCANotFound=  errors.New(fmt.Sprintf("Root CA was not found or defined %s", envTLSRootCA))
	ErrCertNotFoundOrDefined = errors.New(fmt.Sprintf("TLS Cert was not found or defined %s", envTLSCert))
	ErrKeyNotFoundOrDefined = errors.New(fmt.Sprintf("TLS Key was not found or defined %s", envTLSKey))
)
