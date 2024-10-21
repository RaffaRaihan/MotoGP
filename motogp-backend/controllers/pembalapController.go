package controllers

import (
	"motogp-api/models"
	"motogp-api/services"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllPembalap(c *gin.Context) {
	pembalaps, err := services.GetAllPembalap()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pembalaps)
}

func CreatePembalap(c *gin.Context) {
	var pembalap models.Pembalap
	if err := c.ShouldBind(&pembalap); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi apakah tim dengan TimID yang diberikan ada
	_, err := services.GetTimById(pembalap.TimID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid TimID"})
		return
	}

	// Upload file gambar jika ada
	file, err := c.FormFile("gambar")
	if err == nil {
		filename := uuid.New().String() + filepath.Ext(file.Filename)
		filepath := filepath.Join("uploads", filename)

		// Simpan file gambar
		if err := c.SaveUploadedFile(file, filepath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
			return
		}
		pembalap.Gambar = filename
	}

	if err := services.CreatePembalap(pembalap); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pembalap)
}

func UpdatePembalap(c *gin.Context) {
	// Ambil id dari parameter URL dan konversi dari string ke uint
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Ambil pembalap berdasarkan ID
	pembalap, err := services.GetPembalapById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pembalap not found"})
		return
	}

	// Bind data dari request body ke pembalap
	if err := c.ShouldBindJSON(&pembalap); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update pembalap
	if err := services.UpdatePembalap(pembalap); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pembalap)
}

func DeletePembalap(c *gin.Context) {
	// Ambil id dari parameter URL dan konversi dari string ke uint
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam) // Konversi string ke int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Gunakan id yang sudah dikonversi menjadi uint
	if err := services.DeletePembalap(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pembalap deleted"})
}
