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

func (pu *ProductUseCase) InsertProduct(product models.Product) (models.Product, error) {
	productId, err := pu.repository.InsertProduct(product)
	if err != nil {
		return models.Product{}, err
	}
	product.ID = productId
	return product, nil
}
