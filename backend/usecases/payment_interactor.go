package usecases

import "github.com/zyzmoz/DigitalMarket/domain"

type PaymentInteractor struct {
	PaymentRepository PaymentRepository
}

func (pi *PaymentInteractor) FindAll() (payments domain.Payments, err error) {
	payments, err = pi.PaymentRepository.FindAll()

	return
}

func (pi *PaymentInteractor) FindOne(id string) (payment domain.Payment, err error) {
	payment, err = pi.PaymentRepository.FindOne(id)

	return
}

func (pi *PaymentInteractor) Create(paymentData domain.Payment) (payment domain.Payment, err error) {
	payment, err = pi.PaymentRepository.Create(paymentData)

	return
}

func (pi *PaymentInteractor) Update(paymentData domain.Payment) (payment domain.Payment, err error) {
	payment, err = pi.PaymentRepository.Update(paymentData)

	return
}

func (pi *PaymentInteractor) Delete(paymentData domain.Payment) (err error) {
	err = pi.PaymentRepository.Delete(paymentData)

	return
}
