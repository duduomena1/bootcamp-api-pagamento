package controller

import (
	"github.com/gin-gonic/gin"
	"go-api/usecase"
	"net/http"
)

type ProductController struct {
	productUseCase *usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) *ProductController {
	return &ProductController{
		productUseCase: &usecase,
	}
}

func (p *ProductController) GetProducts(c *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)

}
