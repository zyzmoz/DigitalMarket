package usecases

import "github.com/zyzmoz/DigitalMarket/domain"

type ProductInteractor struct {
	ProductRepository ProductRepository
}

func (pi *ProductInteractor) FindAll() (products domain.Products, err error) {
	products, err = pi.ProductRepository.FindAll()

	return
}

func (pi *ProductInteractor) FindOne(id string) (product domain.Product, err error) {
	product, err = pi.ProductRepository.FindOne(id)

	return
}

func (pi *ProductInteractor) Create(productData domain.Product) (product domain.Product, err error) {
	product, err = pi.ProductRepository.Create(productData)

	return
}

func (pi *ProductInteractor) Update(productData domain.Product) (product domain.Product, err error) {
	product, err = pi.ProductRepository.Update(productData)

	return
}

func (pi *ProductInteractor) Delete(productData domain.Product) (err error) {
	err = pi.ProductRepository.Delete(productData)

	return
}
