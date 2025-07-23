package main

import (
	"midtrans-go/controller"
	"midtrans-go/initializer"
	"midtrans-go/middleware"
	"midtrans-go/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gin-contrib/cors"
)

func init() {
	initializer.LoadEnv()
}

func main() {

	validate := validator.New()
	midtransService := service.NewMidtransServiceImpl(validate)
	midtransController := controller.NewMidtransControllerImpl(midtransService)

	r := gin.Default()

	// add CORS
	corsMiddleware := cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	})
	r.Use(corsMiddleware)

	r.Use(middleware.ErrorHandle())
	midtrans := r.Group("/midtrans")
	{
		midtrans.POST("/create", midtransController.Create)
	}

	r.Run()
}