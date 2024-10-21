package services

import (
	"motogp-api/config"
	"motogp-api/models"
)

func GetAllPembalap() ([]models.Pembalap, error) {
	var pembalaps []models.Pembalap
	result := config.GetDB().Find(&pembalaps)
	return pembalaps, result.Error
}

func CreatePembalap(pembalap models.Pembalap) error {
	result := config.GetDB().Create(&pembalap)
	return result.Error
}

func GetPembalapById(id uint) (models.Pembalap, error) {
	var pembalap models.Pembalap
	result := config.GetDB().First(&pembalap, id)
	return pembalap, result.Error
}

func UpdatePembalap(pembalap models.Pembalap) error {
	result := config.GetDB().Save(&pembalap)
	return result.Error
}

func DeletePembalap(id uint) error {
	result := config.GetDB().Delete(&models.Pembalap{}, id)
	return result.Error
}
