package service

import "MyGram/repository"

type Service struct {
	repo repository.RepoInterface
}

type ServiceInterface interface {
	UserInterface
	PhotoInterface
	CommentInterface
	SocialMediaInterface
}

func NewService(repo repository.RepoInterface) ServiceInterface {
	return &Service{repo: repo}
}
