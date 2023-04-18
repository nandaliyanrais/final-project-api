package middlewares

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"mygram-api/models/domain"
	"mygram-api/models/response"
	"mygram-api/photos/service"
)

func Authorization(photoService service.PhotoService) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        var (
            photo domain.Photo
            err   error
        )

        photoID, _ := strconv.Atoi(ctx.Param("id"))
        userData := ctx.MustGet("userData").(jwt.MapClaims)
        userID := uint(userData["id"].(float64))

        if photo, err = photoService.GetOne(uint(photoID)); err != nil {
            ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorResponse{
                Code:   http.StatusBadRequest,
                Status: "Bad Request",
                Errors: gin.H{
                    "message": "Photo not found",
                },
            })

            return
        }

        if photo.UserID != userID {
            ctx.AbortWithStatusJSON(http.StatusForbidden, response.ErrorResponse{
                Code:   http.StatusForbidden,
                Status: "Forbidden",
                Errors: gin.H{
                    "message": "You don't have permission",
                },
            })

            return
        }

        ctx.Next()
    }
}
