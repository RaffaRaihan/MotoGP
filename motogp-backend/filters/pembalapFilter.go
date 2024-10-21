package filters

import (
	"motogp-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FilterPembalap(c *gin.Context) {
	riders := []models.Pembalap{}
	filteredPembalap := make([]map[string]interface{}, len(riders))
	for i, rider := range riders {
		filteredPembalap[i] = map[string]interface{}{
			"nama":   rider.Nama,
			"nomer":  rider.Nomer,
			"negara": rider.Negara,
			"gambar": rider.Gambar,
		}
	}

	c.JSON(http.StatusOK, filteredPembalap)
}