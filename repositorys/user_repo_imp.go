package repository

import (
	"errors"

	"github.com/DavidAfdal/go-gin-rest/helper"
	"github.com/DavidAfdal/go-gin-rest/models"
	"gorm.io/gorm"
)

type UserRepoImpl struct {
	Db *gorm.DB
}


func NewUserRepoImpl(Db *gorm.DB) UserRepository {
	return &UserRepoImpl{Db: Db}
} 

func (u *UserRepoImpl) Delete(userId int) {
	var user models.User
	result := u.Db.Where("user_id = ?", userId).Delete(&user)
	helper.ErrorPanic(result.Error)
}

func (u *UserRepoImpl) FindAll() []models.User {
	var users []models.User
	result := u.Db.Find(&users)
	helper.ErrorPanic(result.Error)
	return users
}

func (u *UserRepoImpl) FindById(userId int) (models.User, error) {
	var user models.User
	result := u.Db.Find(&user, userId)
	if result != nil {
		return user, result.Error
	} else {
		return user, nil
	}
}

func (u *UserRepoImpl) Save(user models.User) {
	result := u.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}

func (u *UserRepoImpl) FindByEmail(email string) (models.User, error) {
	var user models.User
    result := u.Db.First(&user, "email = ?", email)
	if result.Error != nil {
		return user, errors.New("Invalid Email")
	}
	return user, nil
}