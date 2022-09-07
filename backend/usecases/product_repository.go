package usecases

import "github.com/zyzmoz/DigitalMarket/domain"

type ProductRepository interface {
	FindAll() (domain.Products, error)
	FindOne(string) (domain.Product, error)
	Create(domain.Product) (domain.Product, error)
	Update(domain.Product) (domain.Product, error)
	Delete(domain.Product) error
}
