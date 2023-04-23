package service

import "MyGram/domain"

type SocialMediaInterface interface {
	CreateSocialMedia(in domain.SocialMedia) (res domain.SocialMedia, err error)
	GetAllSocialMedia(userId uint) (res []domain.SocialMedia, err error)
	GetSocialMediaById(in domain.SocialMedia) (res domain.SocialMedia, err error)
	UpdateSocialMedia(in domain.SocialMedia) (res domain.SocialMedia, err error)
	DeleteSocialMedia(id int) (err error)
}

func (s *Service) CreateSocialMedia(in domain.SocialMedia) (res domain.SocialMedia, err error) {
	res, err = s.repo.CreateSocialMedia(in)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) GetAllSocialMedia(userId uint) (res []domain.SocialMedia, err error) {
	res, err = s.repo.GetAllSocialMedia(userId)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) GetSocialMediaById(in domain.SocialMedia) (res domain.SocialMedia, err error) {
	res, err = s.repo.GetSocialMediaById(in)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (s *Service) UpdateSocialMedia(in domain.SocialMedia) (res domain.SocialMedia, err error) {
	res, err = s.repo.UpdateSocialMedia(in)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (s *Service) DeleteSocialMedia(id int) (err error) {
	err = s.repo.DeleteSocialMedia(id)
	if err != nil {
		return err
	}
	return nil
}
