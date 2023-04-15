package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username    string        `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required-Your username is required"`
	Email       string        `gorm:"not null;uniqueIndex" json:"email" validate:"required-Email is required,email-Invalid email format"`
	Password    string        `gorm:"not null" json:"password" validate:"required-Password is required,MinStringLength(6)-Password has to have a minimum length of 6 characters"`
	Age         int           `gorm:"not null" json:"age" validate:"required-Age is required,range(8|100)-age has to be above 8 years old"`
	Photo       []Photo       `gorm:"constaint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photo" validate:"required-Photo is required"`
	Comments    []Cpmment     `gorm:"constaint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments" validate:"required-Comments is required"`
	SocialMedia []SocialMedia `gorm:"constaint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_media" validate:"required-Social Media is required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
