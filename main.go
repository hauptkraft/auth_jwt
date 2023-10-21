package main

import (
	"go-jwt/config"
	"go-jwt/controller"
	"go-jwt/helper"
	"go-jwt/model"
	"go-jwt/repository"
	"go-jwt/router"
	"go-jwt/service"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

func main() {
	// loadConfig, err := config.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("зависимости не подтянулись", err)
	// }
	// fmt.Println(loadConfig)
	// router := gin.Default()

	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, "welcome home")
	// })

	// server := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: router,
	// }

	// ser := server.ListenAndServe()
	// helper.ErrorPanic(ser)
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("переменные оружения не загрузились", err)
	}

	//БД
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("users").AutoMigrate(&model.Users{})

	//инициализация репозитория
	userRepository := repository.NewUsersRepositoryImpl(db)

	//Init Service
	authenticationService := service.NewAuthenticationServiceImpl(userRepository, validate)

	//Init controller
	authenticationController := controller.NewAuthenticationController(authenticationService)
	usersController := controller.NewUsersController(userRepository)

	routes := router.NewRouter(userRepository, authenticationController, usersController)

	server := &http.Server{
		Addr:           ":" + loadConfig.ServerPort,
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe()
	helper.ErrorPanic(server_err)

}
