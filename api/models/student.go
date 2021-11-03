package models

import (
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type StudentLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type StudentReg struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	RegNo    string `json:"regno"`
}

type Student struct {
	UUID     uuid.UUID `gorm:"primary_key" json:"uuid"`
	Email    string    `gorm:"size:255;not null;" json:"email"`
	Password string    `gorm:"size:255;not null;" json:"password"`
	Name     string    `gorm:"size:255;not null;" json:"name"`
	RegNo    string    `gorm:"size:255;not null;" json:"regno"`
}

type StudentData struct {
	UUID  uuid.UUID `gorm:"primary_key" json:"uuid"`
	Email string    `gorm:"size:255;not null;" json:"email"`
	Name  string    `gorm:"size:255;not null;" json:"name"`
	RegNo string    `gorm:"size:255;not null;" json:"regno"`
}
