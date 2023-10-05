package main

import (
	"net/http"
	"time"

	"github.com/DavidAfdal/go-gin-rest/config"
	"github.com/DavidAfdal/go-gin-rest/controllers"
	"github.com/DavidAfdal/go-gin-rest/helper"
	"github.com/DavidAfdal/go-gin-rest/models"
	repository "github.com/DavidAfdal/go-gin-rest/repositorys"
	"github.com/DavidAfdal/go-gin-rest/routes"
	"github.com/DavidAfdal/go-gin-rest/services"
	"github.com/go-playground/validator/v10"
)


func main() {
	Init()
}

func Init() {
	db := config.DatabaseConnection()
	db.Table("users").AutoMigrate(&models.User{})
	validate := validator.New()

	userRepo := repository.NewUserRepoImpl(db)
	userServices := services.NewUserServiceImpl(userRepo, validate)
	userController := controllers.NewUserController(userServices)
	route := routes.NewRouter(userController)

	server := &http.Server{
        Addr: ":5000",
		Handler: route,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	serve_err := server.ListenAndServe()
	helper.ErrorPanic(serve_err)
}
