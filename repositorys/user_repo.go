package repository

import "github.com/DavidAfdal/go-gin-rest/models"

type UserRepository interface {
	Save(user models.User)
	Delete(userId int)
	FindById(userId int) (models.User, error)
	FindAll() []models.User
	FindByEmail(email string) (models.User, error)
}