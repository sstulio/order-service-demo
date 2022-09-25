package main

import (
	"log"

	"github.com/Netflix/go-env"
	"github.com/gin-gonic/gin"
	"github.com/sstulio/order-service-demo/internal/app"
	"github.com/sstulio/order-service-demo/internal/pkg/config"
	"github.com/sstulio/order-service-demo/internal/pkg/database"
)

var environment config.Environment

func init() {
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	_, err := database.InitDatabase(environment.DatabaseDNS)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/api/orders/", app.GetOrders)

	if err := r.Run(); err != nil {
		log.Fatalf("could not initiate web server %v", err)
	}

}
