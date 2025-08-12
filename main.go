package main

import (
	"template-go/config"
	"template-go/controllers"
	"template-go/repositories"
	"template-go/services"

	"github.com/gin-gonic/gin"
)

func main() {
	//Database intialisation
	db := config.InitDB()

	//Repository intialisation
	userRepo := repositories.NewUserRepository(db)

	//Service intialisation
	userService := services.NewUserService(userRepo)

	//intialise gin server in go
	server := gin.Default()

	//Controller intialisation
	controllers.SetupUserRoutes(server, userService)

	//Server Start
	server.Run(":80")
}
