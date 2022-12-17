package usecases

import "github.com/zyzmoz/DigitalMarket/domain"

type PaymentRepository interface {
	FindAll() (domain.Payments, error)
	FindOne(string) (domain.Payment, error)
	Create(domain.Payment) (domain.Payment, error)
	Update(domain.Payment) (domain.Payment, error)
	Delete(domain.Payment) error
}
