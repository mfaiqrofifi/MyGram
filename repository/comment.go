package repository

import (
	// "ShowCase/repository/query"
	"MyGram/domain"
	"errors"
	"fmt"

	"gorm.io/gorm/clause"
)

type CommentInterface interface {
	CreateComment(res domain.Comment) (domain.Comment, error)
	GetAllComment(userId uint) (res []domain.Comment, err error)
	GetCommentById(in domain.Comment) (res domain.Comment, err error)
	UpdateComment(in domain.Comment) (res domain.Comment, err error)
	DeleteComment(id int) (err error)
}

func (r Repo) CreateComment(res domain.Comment) (domain.Comment, error) {
	err := r.gorm.Debug().Create(&res).Error
	if err != nil {
		return domain.Comment{}, err
	}
	return res, nil
}

func (r Repo) GetAllComment(userId uint) (res []domain.Comment, err error) {
	err = r.gorm.Debug().Where("user_id=?", userId).Preload(clause.Associations).Find(&res).Error
	if err != nil {
		return []domain.Comment{}, err
	}
	return res, nil
}
func (r Repo) GetCommentById(in domain.Comment) (res domain.Comment, err error) {
	err = r.gorm.Debug().Joins("Photo").Joins("User").First(&res, in.ID).Error
	if err != nil {
		return domain.Comment{}, err
	}
	return res, nil
}

func (r Repo) UpdateComment(in domain.Comment) (res domain.Comment, err error) {
	res = in
	err = r.gorm.Model(&res).Where("id=?", in.ID).Updates(domain.Comment{
		Message: res.Message,
	}).Error
	if err != nil {
		return res, err
	}
	return res, nil
}
func (r Repo) DeleteComment(id int) (err error) {

	comment := domain.Comment{}
	data := r.gorm.Where("id=?", id).Delete(&comment)

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
