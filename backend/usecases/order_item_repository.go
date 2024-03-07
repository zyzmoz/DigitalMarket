package usecases

import "github.com/zyzmoz/DigitalMarket/domain"

type OrderItemRepository interface {
	FindAll() (domain.OrderItems, error)
	FindOne(string) (domain.OrderItem, error)
	Create(domain.OrderItem) (domain.OrderItem, error)
	Update(domain.OrderItem) (domain.OrderItem, error)
	Delete(domain.OrderItem) error
}
