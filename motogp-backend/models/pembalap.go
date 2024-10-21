package models

import (
	"gorm.io/gorm"
)

type Pembalap struct {
	gorm.Model
	Nama    string `json:"nama"`
	Nomer	int16  `json:"nomer"`
	TimID   uint   `json:"tim_id"`   // Foreign key dari Tim
	Tim     Tim    `json:"tim"`      // Relasi ke tabel Tim
	Negara  string `json:"negara"`
	Gambar  string `json:"gambar"`   // Nama file gambar yang di-upload
}
