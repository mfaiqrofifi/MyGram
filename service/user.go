package service

import (
	"MyGram/domain"
)

type UserInterface interface {
	UserRegister(in domain.User) (res domain.User, err error)
	UserLogin(in domain.User) (res domain.User, err error)
}

func (s *Service) UserRegister(in domain.User) (res domain.User, err error) {
	res, err = s.repo.Register(in)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) UserLogin(in domain.User) (res domain.User, err error) {
	res, err = s.repo.UserLogin(in)
	if err != nil {
		return res, err
	}
	return res, nil
}
