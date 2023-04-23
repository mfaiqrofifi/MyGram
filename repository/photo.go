package repository

import (
	// "ShowCase/repository/query"
	"MyGram/domain"
	"errors"
	"fmt"

	"gorm.io/gorm/clause"
)

type PhotoInterface interface {
	CreatePhoto(res domain.Photo) (domain.Photo, error)
	GetAllPhoto(userId uint) (res []domain.Photo, err error)
	GetPhotoById(in domain.Photo) (res domain.Photo, err error)
	UpdatePhoto(in domain.Photo) (res domain.Photo, err error)
	DeletePhoto(id int) (err error)
}

func (r Repo) CreatePhoto(res domain.Photo) (domain.Photo, error) {
	err := r.gorm.Debug().Preload("Users").Create(&res).Error
	if err != nil {
		return domain.Photo{}, err
	}
	return res, nil
}

func (r Repo) GetAllPhoto(userId uint) (res []domain.Photo, err error) {
	err = r.gorm.Debug().Where("user_id=?", userId).Preload(clause.Associations).Find(&res).Error
	if err != nil {
		return []domain.Photo{}, err
	}
	return res, nil
}
func (r Repo) GetPhotoById(in domain.Photo) (res domain.Photo, err error) {
	err = r.gorm.Debug().Joins("User").First(&res, in.ID).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r Repo) UpdatePhoto(in domain.Photo) (res domain.Photo, err error) {
	res = in
	err = r.gorm.Model(&res).Where("id=?", in.ID).Updates(domain.Photo{
		Title:    res.Title,
		Caption:  res.Caption,
		PhotoUrl: res.PhotoUrl,
	}).Error
	if err != nil {
		return res, err
	}
	return res, nil
}
func (r Repo) DeletePhoto(id int) (err error) {

	book := domain.Photo{}
	data := r.gorm.Where("id=?", id).Delete(&book)

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
