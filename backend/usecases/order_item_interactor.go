package usecases

import "github.com/zyzmoz/DigitalMarket/domain"

type OrderItemInteractor struct {
	OrderItemRepository OrderItemRepository
}

func (oi *OrderItemInteractor) FindAll() (orderItems domain.OrderItems, err error) {
	orderItems, err = oi.OrderItemRepository.FindAll()

	return
}

func (oi *OrderItemInteractor) FindOne(id string) (orderItem domain.OrderItem, err error) {
	orderItem, err = oi.OrderItemRepository.FindOne(id)

	return
}

func (oi *OrderItemInteractor) Create(orderItemData domain.OrderItem) (orderItem domain.OrderItem, err error) {
	orderItem, err = oi.OrderItemRepository.Create(orderItemData)

	return
}

func (oi *OrderItemInteractor) Update(orderItemData domain.OrderItem) (orderItem domain.OrderItem, err error) {
	orderItem, err = oi.OrderItemRepository.Update(orderItemData)

	return
}

func (oi *OrderItemInteractor) Delete(orderItemData domain.OrderItem) (err error) {
	err = oi.OrderItemRepository.Delete(orderItemData)

	return
}
