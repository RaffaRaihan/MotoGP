package routes

import (
	"motogp-api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Tambahkan middleware CORS
	router.Use(cors.Default())

	router.Static("/uploads", "./uploads")

	v1 := router.Group("/api/v1")
	{
		// Rute Pembalap
		v1.GET("/pembalap", controllers.GetAllPembalap)
		v1.POST("/pembalap", controllers.CreatePembalap)
		v1.PUT("/pembalap/:id", controllers.UpdatePembalap)
		v1.DELETE("/pembalap/:id", controllers.DeletePembalap)

		// Rute Tim
		v1.GET("/tim", controllers.GetAllTim)
		v1.POST("/tim", controllers.CreateTim)
		v1.PUT("/tim/:id", controllers.UpdateTim)
		v1.DELETE("/tim/:id", controllers.DeleteTim)
	}

	return router
}
