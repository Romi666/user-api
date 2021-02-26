package repo

import (
	"github.com/jinzhu/gorm"
	"user-api/model"
	"user-api/user"
)

type UserRepoImpl struct {
	DB *gorm.DB
}

func (u *UserRepoImpl) Create(user *model.UserData) error {
	err := u.DB.Table("user").Create(model.UserData{NamaUser: user.NamaUser, TL: user.TL, NoKTP: user.NoKTP, Pekerjaan: user.Pekerjaan, PT: user.PT}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepoImpl) Update(userID int, user *model.UserData) error {
	err := u.DB.Table("user").Where("id = ?", userID).Updates(map[string]interface{}{"nama_user" : user.NamaUser, "tanggal_lahir":user.TL, "no_ktp" : user.NoKTP, "pekerjaan":user.Pekerjaan, "pendidikan_terakhir" : user.PT}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepoImpl) GetAll() (*[]model.UserTable, error) {
	var users []model.UserTable
	err := u.DB.Table("user").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (u *UserRepoImpl) GetByID(userID int) (*model.UserTable, error) {
	var user model.UserTable
	err := u.DB.Table("user").Where("id = ?", userID).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUserRepo(DB *gorm.DB) user.UserRepo {
	return &UserRepoImpl{DB}
}
