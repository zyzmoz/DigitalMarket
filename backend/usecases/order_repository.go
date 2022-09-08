package usecases

import "github.com/zyzmoz/DigitalMarket/domain"

type OrderRepository interface {
	FindAll() (domain.Orders, error)
	FindOne(string) (domain.Order, error)
	Create(domain.Order) (domain.Order, error)
	Update(domain.Order) (domain.Order, error)
	Delete(domain.Order) error
}
