package main

import (
	"api/controller"
	"api/db"
	"api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	dbConnection, err := db.DbConnection()

	if err != nil {
		panic(err)
	}

	ProductUseCase := usecase.NewProductUseCase()

	productController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)
	server.Run(":8080")
}
