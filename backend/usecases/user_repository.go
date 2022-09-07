package usecases

import "github.com/zyzmoz/DigitalMarket/domain"

type UserRepository interface {
	FindAll() (domain.Users, error)
	FindOne(string) (domain.User, error)
	Create(domain.User) (domain.User, error)
	Update(domain.User) (domain.User, error)
	Delete(domain.User) error
}
