package models

import (
	"gopg_api/src/database"

	"gorm.io/gorm"
)

type Snus struct {
	gorm.Model
	Brand string `gorm:"size:255;not null;unique" json:"brand"`
	Flavour string `gorm:"size:255;not null;unique" json:"flavour"`
	Price float64 `gorm:"size:255;not null;" json:"price"`
}

func (s *Snus) Save() (*Snus, error) {
	err := database.Database.Create(&s).Error
	if err != nil {
		return &Snus{}, err
	}
	return s, nil
}