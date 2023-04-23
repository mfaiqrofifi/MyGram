package config

import (
	"MyGram/domain"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gorm struct {
	Username string
	Password string
	Port     string
	Address  string
	Database string
	DB       *gorm.DB
}

type GormDb struct {
	*Gorm
}

var (
	GORM *GormDb
)

func InitPostgrees() error {
	GORM = new(GormDb)
	GORM.Gorm = &Gorm{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Address:  os.Getenv("POSTGRES_ADDRESS"),
		Database: os.Getenv("POSTGRES_DB"),
	}
	err := GORM.Gorm.OpenConnection()
	if err != nil {
		return err
	}
	// err = PSQL.DB.Ping()
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (p *Gorm) OpenConnection() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", p.Address, p.Port, p.Username, p.Password, p.Database)
	dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	p.DB = dbConnection
	err = p.DB.Debug().AutoMigrate(domain.User{}, domain.Photo{}, domain.Comment{}, domain.SocialMedia{})
	if err != nil {
		return err
	}
	fmt.Println("Succes to Connect using GORM")
	return nil
}
