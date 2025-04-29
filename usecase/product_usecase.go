package usecase

import (
	"go-api/models"
	"go-api/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}

func (pu *ProductUseCase) GetProducts() ([]models.Product, error) {

	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) InsertProduct(product models.Product) (int, error) {

	return pu.repository.InsertProduct(product)
}
