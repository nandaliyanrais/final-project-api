package routes

import (
	"github.com/gin-gonic/gin"

	"mygram-api/database"
	"mygram-api/users/controller"
	"mygram-api/users/repository"
	"mygram-api/users/service"
)

func UserRoute(router *gin.Engine) {
	
	db := database.StartDB()

	repositoryUser := repository.NewUserRepository(db)
	serviceUser := service.NewUserService(repositoryUser)
	controllerUser := controller.NewUserController(serviceUser)

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controllerUser.Register)
		userRouter.POST("/login", controllerUser.Login)
	}

}