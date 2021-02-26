package user

import "user-api/model"

type UserUsecase interface {
	Create(user *model.UserData) error
	Update(userID int, user*model.UserData) error
	GetAll() (*[]model.UserTable, error)
	GetByID(userID int) (*model.UserTable, error)
}