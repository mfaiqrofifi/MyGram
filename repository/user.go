package repository

import (
	// "ShowCase/repository/query"
	"MyGram/domain"
)

type UserInterface interface {
	Register(in domain.User) (domain.User, error)
	UserLogin(in domain.User) (domain.User, error)
}

func (r Repo) Register(res domain.User) (domain.User, error) {
	err := r.gorm.Debug().Create(&res).Error
	if err != nil {
		return domain.User{}, err
	}
	return res, nil
}

func (r Repo) UserLogin(res domain.User) (domain.User, error) {
	err := r.gorm.Debug().Where("email=?", res.Email).Take(&res).Error
	if err != nil {
		return domain.User{}, err
	}
	return res, nil
}
