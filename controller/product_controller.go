package controller

import (
	"github.com/gin-gonic/gin"
	"go-api/models"
	"go-api/usecase"
	"net/http"
	"strconv"
)

type ProductController struct {
	productUseCase *usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) *ProductController {
	return &ProductController{
		productUseCase: &usecase,
	}
}

// GetProducts godoc
// @Summary Get all products
// @Description Retrieve a list of all products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
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

// InsertProduct godoc
// @Summary Insert a new product
// @Description Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product object"
// @Success 201 {object} models.Product
// @Router /product [post]
func (p *ProductController) InsertProduct(c *gin.Context) {
	var product models.Product
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	insertedProduct, err := p.productUseCase.InsertProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, insertedProduct)
}

// GetProductsByID godoc
// @Summary Get product by ID
// @Description Get product details by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Failure 400 {object} models.Response
// @Failure 404 {object} models.Response
// @Router /product/{id} [get]
func (p *ProductController) GetProductsByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response := models.Response{
			Message: "ID is required",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	Id, err := strconv.Atoi(id)
	if err != nil {
		response := models.Response{
			Message: "ID must be a number",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	product, err := p.productUseCase.GetProductById(Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Product not found",
		})
		return
	}
	c.JSON(http.StatusOK, product)

}
