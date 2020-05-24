package service

import (
	"github.com/stevepartridge/blogorama/pkg/certs"
	"github.com/stevepartridge/blogorama/pkg/log"
)

// TLS sets up the public/private and root cert authority
func (s *Service) TLS() {

	rootCA, err := certs.GetRootCA()
	log.IfError(err)
	if rootCA != nil {
		err = s.Server.AppendCertsFromPEM(rootCA)
		log.IfError(err)
	}

	cert, err := certs.GetCert()
	log.IfError(err)
	key, err := certs.GetKey()
	log.IfError(err)

	err = s.Server.AddKeyPair(cert, key)
	log.IfError(err)

}
