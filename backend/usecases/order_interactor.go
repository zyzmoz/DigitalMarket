package usecases

import "github.com/zyzmoz/DigitalMarket/domain"

type OrderInteractor struct {
	OrderRepository OrderRepository
}

func (oi *OrderInteractor) FindAll() (orders domain.Orders, err error) {
	orders, err = oi.OrderRepository.FindAll()

	return
}

func (oi *OrderInteractor) FindOne(id string) (order domain.Order, err error) {
	order, err = oi.OrderRepository.FindOne(id)

	return
}

func (oi *OrderInteractor) Create(orderData domain.Order) (order domain.Order, err error) {
	order, err = oi.OrderRepository.Create(orderData)

	return
}

func (oi *OrderInteractor) Update(orderData domain.Order) (order domain.Order, err error) {
	order, err = oi.OrderRepository.Update(orderData)

	return
}

func (oi *OrderInteractor) Delete(orderData domain.Order) (err error) {
	err = oi.OrderRepository.Delete(orderData)

	return
}
