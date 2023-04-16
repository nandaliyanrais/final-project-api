package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"mygram-api/helpers"
	"mygram-api/models/response"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helpers.VerifyToken(ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Errors: err.Error(),
			})

			return
		}

		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}