package domain

import (
	"time"

	"github.com/ajalck/spine/pkg/utils"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint      `json:"-" gorm:"primaryKey"`
	ID26      string    `json:"id26" gorm:"unique;not null;size:36"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

// hook to generate and set id26 before saving the table
func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID26 = utils.GenerateUniqueString()

	return nil
}

type User struct {
	BaseModel
	FullName string `json:"fullname" gorm:"not null"`
	Email    string `json:"email" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}
