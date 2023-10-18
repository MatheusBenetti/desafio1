package usecase

import "github.com/MatheusBenetti/desafio1/domain/model"

type ProductUseCase struct {
	ProductRepository model.ProductRepositoryInterface
}

func (p *ProductUseCase) CreateProduct(name string, description string, price float64) (*model.Product, error) {
	product, err := model.NewProduct(name, description, price)
	if err != nil {
		return nil, err
	}

	err = p.ProductRepository.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductUseCase) FindProducts() ([]model.Product, error) {
	products, err := p.ProductRepository.FindProducts()
	if err != nil {
		return nil, err
	}

	return []model.Product{*products}, nil
}
