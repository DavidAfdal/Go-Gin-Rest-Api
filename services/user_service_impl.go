package services

import (
	"github.com/DavidAfdal/go-gin-rest/data/request"
	"github.com/DavidAfdal/go-gin-rest/data/response"
	"github.com/DavidAfdal/go-gin-rest/helper"
	"github.com/DavidAfdal/go-gin-rest/models"
	repository "github.com/DavidAfdal/go-gin-rest/repositorys"
	"github.com/DavidAfdal/go-gin-rest/utils"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{UserRepository: userRepository, Validate: validate}
}


func (u *UserServiceImpl) Login(user request.LoginRequest) {
	
}
func (u *UserServiceImpl) Register(user request.CreateUserRequest) {
	hashedPassword, err := utils.HashPassword(user.Password)
	helper.ErrorPanic(err)
	newUser := models.User{
		UserName: user.Username,
		Email: user.Email,
		Password: hashedPassword,
	}
	u.UserRepository.Save(newUser)
}

func (u *UserServiceImpl) FindAll() []response.UserResponse {
	result := u.FindAll()
	var users []response.UserResponse
    for _, value := range result {
		user := response.UserResponse{
			Id: value.Id ,
			Username: value.Username,
			Email: value.Email,
		}
		users = append(users, user)
	}
	return users
}

func (u *UserServiceImpl) FindById(userId int) response.UserResponse {
   userData, err := u.UserRepository.FindById(userId)
   helper.ErrorPanic(err)
   userResponse := response.UserResponse{
	Id: userData.Id,
	Username: userData.UserName,
	Email: userData.Email,
   }
   return userResponse
}
