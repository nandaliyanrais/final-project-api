package routes

import (
	"github.com/gin-gonic/gin"

	"mygram-api/database"
	"mygram-api/photos/controller"
	"mygram-api/photos/middlewares"
	"mygram-api/photos/repository"
	"mygram-api/photos/service"
)

func PhotoRoute(router *gin.Engine) {

	db := database.StartDB()

	repositoryPhoto := repository.NewPhotoRepository(db)
	servicePhoto := service.NewPhotoService(repositoryPhoto)
	controllerPhoto := controller.NewPhotoController(servicePhoto)

	photoRouter := router.Group("/photos", middlewares.Authentication())
	{
		photoRouter.POST("/", controllerPhoto.Create)
		photoRouter.GET("/", controllerPhoto.GetAll)
		photoRouter.GET("/:id", controllerPhoto.GetOne)
		photoRouter.PUT("/:id", middlewares.Authorization(servicePhoto), controllerPhoto.Update)
		photoRouter.DELETE("/:id", middlewares.Authorization(servicePhoto), controllerPhoto.Delete)
	}

}