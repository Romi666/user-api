package usecase

import (
	"user-api/model"
	"user-api/user"
)

type UserUsecaseImpl struct {
	userRepo user.UserRepo
}

func (u *UserUsecaseImpl) Create(user *model.UserData) error {
	return u.userRepo.Create(user)
}

func (u *UserUsecaseImpl) Update(userID int, user *model.UserData) error {
	return u.userRepo.Update(userID, user)
}

func (u *UserUsecaseImpl) GetAll() (*[]model.UserTable, error) {
	return u.userRepo.GetAll()
}

func (u *UserUsecaseImpl) GetByID(userID int) (*model.UserTable, error) {
	return u.userRepo.GetByID(userID)
}

func CreateUserUsecase(userRepo user.UserRepo) user.UserUsecase {
	return &UserUsecaseImpl{userRepo}
}
