package interfaces

import (
	"github.com/zyzmoz/DigitalMarket/domain"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func (or *OrderRepository) FindAll() (orders domain.Orders, err error) {
	or.DB.Find(&orders)

	return
}

func (or *OrderRepository) FindOne(id string) (order domain.Order, err error) {
	or.DB.Find(&order, id)

	return
}

func (or *OrderRepository) Create(orderData domain.Order) (order domain.Order, err error) {
	or.DB.Create(&orderData)

	order = orderData
	return
}

func (or *OrderRepository) Update(orderData domain.Order) (order domain.Order, err error) {
	or.DB.First(&order, "ID = ?", orderData.ID)

	or.DB.Model(&order).Save(&orderData)

	return orderData, err
}

func (or *OrderRepository) Delete(orderData domain.Order) (err error) {
	or.DB.Delete(&orderData)

	return
}
