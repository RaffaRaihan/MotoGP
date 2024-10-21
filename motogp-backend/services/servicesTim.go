package services

import (
	"motogp-api/models"
	"motogp-api/config"
)

func GetAllTim() ([]models.Tim, error) {
	var tims []models.Tim
	result := config.GetDB().Preload("Pembalap").Find(&tims)
	return tims, result.Error
}

func CreateTim(tim models.Tim) error {
	result := config.GetDB().Create(&tim)
	return result.Error
}

func GetTimById(id uint) (models.Tim, error) {
	var tim models.Tim
	result := config.GetDB().Preload("Pembalap").First(&tim, id)
	return tim, result.Error
}

func UpdateTim(tim models.Tim) error {
	result := config.GetDB().Save(&tim)
	return result.Error
}

func DeleteTim(id uint) error {
	result := config.GetDB().Delete(&models.Tim{}, id)
	return result.Error
}
