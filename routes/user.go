package routes

import (
	"github.com/DavidAfdal/go-gin-rest/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter(userController *controllers.UserController) *gin.Engine{
	service := gin.Default()

	router := service.Group("/api")
	router.POST("/register", userController.Register)


	return service
}