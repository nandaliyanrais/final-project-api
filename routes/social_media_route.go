package routes

import (
	"github.com/gin-gonic/gin"

	"mygram-api/database"
	"mygram-api/social_medias/controller"
	"mygram-api/social_medias/middlewares"
	"mygram-api/social_medias/repository"
	"mygram-api/social_medias/service"
)

func SocialMediaRoute(router *gin.Engine) {

	db := database.StartDB()

	repositorySocialMedia := repository.NewSocialMediaRepository(db)
	serviceSoacialMedia := service.NewSocialMediaService(repositorySocialMedia)
	controllerSocialMedia := controller.NewSocialMediaController(serviceSoacialMedia)

	socialMedia := router.Group("/social-media", middlewares.Authentication())
	{
		socialMedia.GET("/", controllerSocialMedia.GetAll)
		socialMedia.GET("/:id", controllerSocialMedia.GetOne)
		socialMedia.POST("/", controllerSocialMedia.Create)
		socialMedia.PUT("/:id", middlewares.Authorization(serviceSoacialMedia), controllerSocialMedia.Update)
		socialMedia.DELETE("/:id", middlewares.Authorization(serviceSoacialMedia), controllerSocialMedia.Delete)
	}

}