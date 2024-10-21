package controllers

import (
	"motogp-api/models"
	"motogp-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTim(c *gin.Context) {
	tims, err := services.GetAllTim()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tims)
}

func CreateTim(c *gin.Context) {
	var tim models.Tim
	if err := c.ShouldBindJSON(&tim); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateTim(tim); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tim)
}

func UpdateTim(c *gin.Context) {
	// Ambil id dari parameter URL dan konversi dari string ke uint
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam) // Konversi string ke int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Gunakan id yang sudah dikonversi
	tim, err := services.GetTimById(uint(id)) // Konversi dari int ke uint
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tim not found"})
		return
	}

	if err := c.ShouldBindJSON(&tim); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateTim(tim); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tim)
}

func DeleteTim(c *gin.Context) {
	// Ambil id dari parameter URL dan konversi dari string ke uint
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam) // Konversi string ke int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Gunakan id yang sudah dikonversi
	if err := services.DeleteTim(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tim deleted"})
}