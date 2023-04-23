package repository

import (
	// "database/sql"

	"gorm.io/gorm"
)

type Repo struct {
	// db   *sql.DB
	gorm *gorm.DB
}

type RepoInterface interface {
	UserInterface
	PhotoInterface
	CommentInterface
	SocialMediaInterface
}

func NewRepo(gorm *gorm.DB) *Repo {
	return &Repo{gorm: gorm}
}
