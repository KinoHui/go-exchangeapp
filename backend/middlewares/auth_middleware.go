package middlewares

import (
	"exchangeapp/backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"errpo": "Missing Authorization Header"})
			ctx.Abort()
		}

		username, err := utils.ParseJWT(token)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"errpo": "Invalid token"})
			ctx.Abort()
		}

		ctx.Set("username", username)
		ctx.Next()
	}
}
