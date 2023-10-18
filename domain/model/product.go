package model

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"time"
)

type ProductRepositoryInterface interface {
	CreateProduct(product *Product) error
	FindProducts() (*Product, error)
}

type Product struct {
	ID          uint16    `json:"id" gorm:"type:autoIncrement;primary_key" valid:"-"`
	Name        string    `json:"name" valid:"required"`
	Description string    `json:"description" valid:"required"`
	Price       float64   `json:"price" valid:"required"`
	CreatedAt   time.Time `json:"created_at" valid:"-"`
}

func (p *Product) isValid() error {
	_, err := govalidator.ValidateStruct(p)

	if p.Price < 0 {
		err := errors.New("price must be greater than 0")
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func generateProductID() uint16 {
	var productIDCounter uint16
	productIDCounter++
	return productIDCounter
}

func NewProduct(name string, description string, price float64) (*Product, error) {
	product := Product{
		Name:        name,
		Description: description,
		Price:       price,
	}

	product.ID = generateProductID()
	product.CreatedAt = time.Now()

	if err := product.isValid(); err != nil {
		return nil, err
	}

	return &product, nil
}
