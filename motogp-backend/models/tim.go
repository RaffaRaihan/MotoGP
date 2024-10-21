package models

import (
	"gorm.io/gorm"
)

type Tim struct {
	gorm.Model
	Nama     string      `json:"nama"`
	Pembalap []Pembalap  `json:"pembalap"` // Relasi One-to-Many, satu tim memiliki banyak pembalap
}
