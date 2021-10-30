package models

import (
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Student struct {
	UUID     uuid.UUID `gorm:"primary_key" json:"uuid"`
	Email    string    `gorm:"size:255;not null;" json:"email"`
	Password string    `gorm:"size:255;not null;" json:"password"`
	Name     string    `gorm:"size:255;not null;" json:"name"`
	RegNo    string    `gorm:"size:255;not null;" json:"regno"`
}
