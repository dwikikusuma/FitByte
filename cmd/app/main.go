package main

import (
	"FitByte/configs"
	"FitByte/internal/handlers"
	"FitByte/internal/infrastructure"
	"FitByte/internal/repositories"
	"FitByte/internal/service"

	"github.com/gin-gonic/gin"

	"FitByte/pkg/log"
)

func main() {
	r := gin.Default()

	log.InitLogger()

	appConfig := configs.LoadConfig(
		configs.WithConfigFolder([]string{"./configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	db := infrastructure.InitDB(appConfig)

	userRepo := repositories.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(r, userService)
	userHandler.SetupRoutes()

	_ = r.Run(":" + appConfig.App.Port)
}
