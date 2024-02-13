package domain

import (
	"github.com/ajalck/spine/pkg/utils"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID   *gorm.Model
	ID26 string `json:"id26" gorm:"unique;not null;size:16"`
}

// hook to generate and set id26 before saving the table
func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID26 = utils.GenerateUniqueString()

	return nil
}

type User struct {
	Id       string `gorm:"primary key"`
	FullName string `json:"fullname" gorm:"not null"`
	Email    string `json:"email" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}
