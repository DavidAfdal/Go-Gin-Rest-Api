package services

import (
	"github.com/DavidAfdal/go-gin-rest/data/request"
	"github.com/DavidAfdal/go-gin-rest/data/response"
)

type UserService interface {
	Login(user request.LoginRequest )
	Register(user request.CreateUserRequest )
	FindAll() []response.UserResponse
	FindById(id int) response.UserResponse
}