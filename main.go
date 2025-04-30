// @title Bootcamp Payment API
// @version 1.0
// @description This is a RESTful API for managing products.
// @host localhost:8000
// @BasePath /
// @schemes http

package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-api/controller"
	"go-api/db"
	_ "go-api/docs"
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
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.Run(":8000")
}
