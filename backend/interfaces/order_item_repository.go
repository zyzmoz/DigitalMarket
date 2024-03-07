package interfaces

import (
	"github.com/zyzmoz/DigitalMarket/domain"
	"gorm.io/gorm"
)

type OrderItemRepository struct {
	DB *gorm.DB
}

func (r *OrderItemRepository) FindAll() (orderItems domain.OrderItems, err error) {
	r.DB.Find(&orderItems)

	return
}

func (r *OrderItemRepository) FindOne(id string) (orderItem domain.OrderItem, err error) {
	r.DB.Find(&orderItem, id)

	return
}

func (r *OrderItemRepository) Create(orderItemData domain.OrderItem) (orderItem domain.OrderItem, err error) {
	r.DB.Create(&orderItemData)

	orderItem = orderItemData

	return
}

func (r *OrderItemRepository) Update(orderItemData domain.OrderItem) (orderItem domain.OrderItem, err error) {
	r.DB.First(&orderItem, "ID = ?", orderItemData.ID)

	r.DB.Model(&orderItem).Save(&orderItemData)

	return
}

func (r *OrderItemRepository) Delete(orderItemData domain.OrderItem) (err error) {
	r.DB.Delete(&orderItemData)

	return
}
