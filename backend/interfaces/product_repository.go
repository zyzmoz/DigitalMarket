package interfaces

import (
	"github.com/zyzmoz/DigitalMarket/domain"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func (pr *ProductRepository) FindAll() (products domain.Products, err error) {
	pr.DB.Find(&products)

	return
}

func (pr *ProductRepository) FindOne(id string) (product domain.Product, err error) {
	pr.DB.Find(&product, id)

	return
}

func (pr *ProductRepository) Create(productData domain.Product) (product domain.Product, err error) {
	pr.DB.Create(&productData)

	product = productData
	return
}

func (pr *ProductRepository) Update(productData domain.Product) (product domain.Product, err error) {
	pr.DB.First(&product, productData.ID)
	pr.DB.Model(&product).Save(&productData)

	return productData, err
}

func (pr *ProductRepository) Delete(productData domain.Product) (err error) {
	pr.DB.Delete(&productData)

	return
}
