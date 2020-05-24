package service

import "github.com/stevepartridge/blogorama/pkg/fauxauth"

func (s *Service) fauxAuth() error {
	var err error
	s.FauxAuth, err = fauxauth.New()
	if err != nil {
		return err
	}

	return nil
}
