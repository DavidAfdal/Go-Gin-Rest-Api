package controllers

import (
	"net/http"

	"github.com/DavidAfdal/go-gin-rest/data/request"
	"github.com/DavidAfdal/go-gin-rest/data/response"
	"github.com/DavidAfdal/go-gin-rest/helper"
	"github.com/DavidAfdal/go-gin-rest/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) *UserController{
	return &UserController{UserService: userService}
}

func (controller *UserController) Login(ctx *gin.Context){
	var loginRequest request.LoginRequest
	err := ctx.BindJSON(&loginRequest)
	helper.ErrorPanic(err)

}

func (controller *UserController) Register(ctx *gin.Context) {
	var createRequest request.CreateUserRequest
	err := ctx.BindJSON(&createRequest)
	helper.ErrorPanic(err)

	controller.UserService.Register(createRequest)
	webResponse := response.WebhookResponse{
		Code: 201,
		Status: "Created",
		Message: "Sucess created user",
		Data: nil,
	}
	ctx.JSON(http.StatusCreated, webResponse)
}