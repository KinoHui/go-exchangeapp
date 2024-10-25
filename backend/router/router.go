package router

import (
	"exchangeapp/backend/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
		auth.GET("/testbook/:id", func(ctx *gin.Context) {
			bookId := ctx.Param("id")
			ctx.JSON(http.StatusOK, gin.H{"book_id": bookId})
		})
	}

	return r
}