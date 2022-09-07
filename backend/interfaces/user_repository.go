package interfaces

import (
	"github.com/jinzhu/gorm"
	"github.com/zyzmoz/DigitalMarket/domain"
)

type UserRepository struct {
	DB *gorm.DB
}

func (ur *UserRepository) FindAll() (users domain.Users, err error) {
	ur.DB.Find(&users)

	return
}

func (ur *UserRepository) FindOne(id string) (user domain.User, err error) {
	ur.DB.Find(&user, id)

	return
}

func (ur *UserRepository) Create(userData domain.User) (user domain.User, err error) {
	ur.DB.Create(&userData)

	return
}

func (ur *UserRepository) Update(userData domain.User) (user domain.User, err error) {
	ur.DB.First(&user, userData.ID)
	ur.DB.Model(&user).Save(&userData)

	return userData, err
}

func (ur *UserRepository) Delete(userData domain.User) (err error) {
	ur.DB.Delete(&userData)

	return
}
