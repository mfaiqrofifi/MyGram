package domain

import (
	"MyGram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string `form:"username" json:"username" gorm:"not null;unique" valid:"required~Your username is required"`
	Email    string `form:"email" json:"email" gorm:"not null;unique" valid:"required~Your email is required,email~Invalid email format"`
	Password string `form:"password" json:"password" gorm:"not null" valid:"required~Your password is required,minstringlength(6)~Password has to have minimum length of 6 characters"`
	Age      int    `form:"age" json:"age" validate:"required~Your age is required,numeric,gte=8~age must be greather than 8"`

	SocialMedia SocialMedia `gorm:"constraint:onUpdate:CASCADE,onDelete:SET NULL;" json:"social_media"`
	Photo   []Photo   `gorm:"constraint:onUpdate:CASCADE,onDelete:SET NULL;" json:"photo"`
	Comment []Comment `gorm:"constraint:onUpdate:CASCADE,onDelete:SET NULL;" json:"comment"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}
	u.Password = helpers.HassPass(u.Password)
	err = nil
	return
}
