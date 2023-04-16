package routes

import (
	"github.com/gin-gonic/gin"

	"mygram-api/database"
	"mygram-api/comments/controller"
	"mygram-api/comments/middlewares"
	"mygram-api/comments/repository"
	photoRepository "mygram-api/photos/repository"
	photoservice "mygram-api/photos/service"
	"mygram-api/comments/service"
)


func CommentRoute(router *gin.Engine) {

	db := database.StartDB()

	repositoryPhoto := photoRepository.NewPhotoRepository(db)
	servicePhoto := photoservice.PhotoService(repositoryPhoto)

	repositoryComment := repository.NewCommentRepository(db)
	serviceComment := service.NewCommentService(repositoryComment)

	controllerComment := controller.NewCommentController(serviceComment, servicePhoto)

	commentRouter := router.Group("/comments", middlewares.Authentication())
	{
		commentRouter.POST("/", controllerComment.Create)
		commentRouter.GET("/", controllerComment.GetAll)
		commentRouter.GET("/:commentId", middlewares.Authorization(serviceComment), controllerComment.GetOne)
		commentRouter.PUT("/:commentId", middlewares.Authorization(serviceComment), controllerComment.Update)
		commentRouter.DELETE("/:commentId", middlewares.Authorization(serviceComment), controllerComment.Delete)
	}

}