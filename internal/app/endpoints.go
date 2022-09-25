package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Order struct {
	ID       string
	Name     string
	Quantity int
}

var orders = []Order{
	{
		ID:       "aasdasd",
		Name:     "Bolsa Zara",
		Quantity: 2,
	},
}

func GetOrders(c *gin.Context) {

	c.JSON(http.StatusOK, orders)
}
