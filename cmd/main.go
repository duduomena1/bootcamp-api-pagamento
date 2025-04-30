package main

import (
	"github.com/gin-gonic/gin"
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	productController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.GET("/products", productController.GetProducts)
	server.POST("/product", productController.InsertProduct)
	server.GET("/product/:id", productController.GetProductsByID)

	server.Run(":8000")
}
