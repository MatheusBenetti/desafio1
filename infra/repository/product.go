package repository

import (
	"github.com/MatheusBenetti/desafio1/domain/model"
	"gorm.io/gorm"
)

type ProductRepositoryDb struct {
	Db *gorm.DB
}

func (p ProductRepositoryDb) CreateProduct(product *model.Product) error {
	err := p.Db.Create(product).Error
	if err != nil {
		return err
	}
	return nil
}

func (p ProductRepositoryDb) FindProducts() ([]model.Product, error) {
	var products []model.Product
	err := p.Db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
