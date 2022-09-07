package usecases

import "github.com/zyzmoz/DigitalMarket/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (ui *UserInteractor) FindAll() (users domain.Users, err error) {
	users, err = ui.UserRepository.FindAll()

	return
}

func (ui *UserInteractor) FindOne(id string) (user domain.User, err error) {
	user, err = ui.UserRepository.FindOne(id)

	return
}

func (ui *UserInteractor) Create(userData domain.User) (user domain.User, err error) {
	user, err = ui.UserRepository.Create(userData)

	return
}

func (ui *UserInteractor) Update(userData domain.User) (user domain.User, err error) {
	user, err = ui.UserRepository.Update(userData)

	return
}

func (ui *UserInteractor) Delete(userData domain.User) (err error) {
	err = ui.UserRepository.Delete(userData)

	return
}
