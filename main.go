package main

import (
	"github.com/gin-gonic/gin"

	_ "mygram-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"mygram-api/routes"
)

// @title MyGram API
// @version 1.0
// @description MyGram is a social media application.
// @termsOfService http://swagger.io/terms/
// @contact.name rais
// @contact.email nandaliyan.rais@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Authorization header using the Bearer scheme
func main() {

	router := gin.Default()

	// Mount Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Set up routes
	routes.UserRoute(router)
	routes.PhotoRoute(router)
	routes.CommentRoute(router)
	routes.SocialMediaRoute(router)
	
	router.Run(":8080")

}
