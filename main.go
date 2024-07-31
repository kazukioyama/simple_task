package main

import (
	"go-rest-api-udemy/controller"
	"go-rest-api-udemy/db"
	"go-rest-api-udemy/repository"
	"go-rest-api-udemy/router"
	"go-rest-api-udemy/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
