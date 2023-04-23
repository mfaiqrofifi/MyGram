package service

import "MyGram/domain"

type PhotoInterface interface {
	CreatePhoto(in domain.Photo) (res domain.Photo, err error)
	GetAllPhoto(userId uint) (res []domain.Photo, err error)
	GetPhotoById(in domain.Photo) (res domain.Photo, err error)
	UpdatePhoto(in domain.Photo) (res domain.Photo, err error)
	DeletePhoto(id int) (err error)
}

func (s *Service) CreatePhoto(in domain.Photo) (res domain.Photo, err error) {
	res, err = s.repo.CreatePhoto(in)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) GetAllPhoto(userId uint) (res []domain.Photo, err error) {
	res, err = s.repo.GetAllPhoto(userId)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) GetPhotoById(in domain.Photo) (res domain.Photo, err error) {
	res, err = s.repo.GetPhotoById(in)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (s *Service) UpdatePhoto(in domain.Photo) (res domain.Photo, err error) {
	res, err = s.repo.UpdatePhoto(in)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (s *Service) DeletePhoto(id int) (err error) {
	err = s.repo.DeletePhoto(id)
	if err != nil {
		return err
	}
	return nil
}
