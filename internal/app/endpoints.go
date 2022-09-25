package app

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sstulio/order-service-demo/internal/models"
	"gorm.io/gorm"
)

func GetOrders(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var orders []models.Order

		err := db.Model(&models.Order{}).Find(&orders).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, "error retrieving orders")
			return
		}

		c.JSON(http.StatusOK, orders)
	}
}

func CreateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer c.Request.Body.Close()

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, "error reading body")
			return
		}

		var order models.Order
		err = json.Unmarshal(body, &order)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, "error parsing order request")
			return
		}

		err = db.Model(&models.Order{}).Save(&order).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, "error creating order")
			return
		}

		c.JSON(http.StatusOK, order)
	}
}
