package service

import "MyGram/domain"

type CommentInterface interface {
	CreateComment(in domain.Comment) (res domain.Comment, err error)
	GetAllComment(userId uint) (res []domain.Comment, err error)
	GetCommentById(in domain.Comment) (res domain.Comment, err error)
	UpdateComment(in domain.Comment) (res domain.Comment, err error)
	DeleteComment(id int) (err error)
}

func (s *Service) CreateComment(in domain.Comment) (res domain.Comment, err error) {
	res, err = s.repo.CreateComment(in)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) GetAllComment(userId uint) (res []domain.Comment, err error) {
	res, err = s.repo.GetAllComment(userId)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) GetCommentById(in domain.Comment) (res domain.Comment, err error) {
	res, err = s.repo.GetCommentById(in)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (s *Service) UpdateComment(in domain.Comment) (res domain.Comment, err error) {
	res, err = s.repo.UpdateComment(in)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (s *Service) DeleteComment(id int) (err error) {
	err = s.repo.DeleteComment(id)
	if err != nil {
		return err
	}
	return nil
}
