package interfaces

import (
	"github.com/zyzmoz/DigitalMarket/domain"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	DB *gorm.DB
}

func (pr *PaymentRepository) FindAll() (Payments domain.Payments, err error) {
	pr.DB.Find(&Payments)

	return
}

func (pr *PaymentRepository) FindOne(id string) (payment domain.Payment, err error) {
	pr.DB.Find(&payment, id)

	return
}

func (pr *PaymentRepository) Create(paymentData domain.Payment) (payment domain.Payment, err error) {
	pr.DB.Create(&paymentData)

	payment = paymentData
	return
}

func (pr *PaymentRepository) Update(paymentData domain.Payment) (payment domain.Payment, err error) {
	pr.DB.First(&payment, "ID = ?", paymentData.ID)
	pr.DB.Model(&payment).Save(&paymentData)

	return paymentData, err
}

func (pr *PaymentRepository) Delete(paymentData domain.Payment) (err error) {
	pr.DB.Delete(&paymentData)

	return
}
