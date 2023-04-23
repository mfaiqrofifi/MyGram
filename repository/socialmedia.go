package repository

import (
	// "ShowCase/repository/query"
	"MyGram/domain"
	"errors"
	"fmt"

	"gorm.io/gorm/clause"
)

type SocialMediaInterface interface {
	CreateSocialMedia(res domain.SocialMedia) (domain.SocialMedia, error)
	GetAllSocialMedia(userId uint) (res []domain.SocialMedia, err error)
	GetSocialMediaById(in domain.SocialMedia) (res domain.SocialMedia, err error)
	UpdateSocialMedia(in domain.SocialMedia) (res domain.SocialMedia, err error)
	DeleteSocialMedia(id int) (err error)
}

func (r Repo) CreateSocialMedia(res domain.SocialMedia) (domain.SocialMedia, error) {
	err := r.gorm.Debug().Create(&res).Error
	if err != nil {
		return domain.SocialMedia{}, err
	}
	return res, nil
}

func (r Repo) GetAllSocialMedia(userId uint) (res []domain.SocialMedia, err error) {
	err = r.gorm.Debug().Where("user_id=?", userId).Preload(clause.Associations).Find(&res).Error
	if err != nil {
		return []domain.SocialMedia{}, err
	}
	return res, nil
}
func (r Repo) GetSocialMediaById(in domain.SocialMedia) (res domain.SocialMedia, err error) {
	err = r.gorm.Debug().Joins("User").First(&res, in.ID).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r Repo) UpdateSocialMedia(in domain.SocialMedia) (res domain.SocialMedia, err error) {
	res = in
	err = r.gorm.Model(&res).Where("id=?", in.ID).Updates(domain.SocialMedia{
		Name:           res.Name,
		SocialMediaUrl: res.SocialMediaUrl,
	}).Error
	if err != nil {
		return res, err
	}
	return res, nil
}
func (r Repo) DeleteSocialMedia(id int) (err error) {

	socialMedia := domain.SocialMedia{}
	data := r.gorm.Where("id=?", id).Delete(&socialMedia)

	fmt.Println(data.RowsAffected)
	err = data.Error
	if err != nil {
		return err
	}
	if data.RowsAffected < 1 {
		return errors.New("Not Found")
	}
	return nil
}
