package router

import (
	"exchangeapp/backend/controllers"
	"exchangeapp/backend/middlewares"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	// cors 配置
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
		auth.GET("/testbook/:id", func(ctx *gin.Context) {
			bookId := ctx.Param("id")
			ctx.JSON(http.StatusOK, gin.H{"book_id": bookId})
		})
	}

	api := r.Group("/api")
	api.GET("/exchangerate", controllers.GetExchangeRates)
	api.Use(middlewares.AuthMiddleWare())
	{
		api.POST("/exchangerate", controllers.CreatExchangeRate)

		api.POST("/article", controllers.CreateArticle)
		api.GET("/article", controllers.GetArticles)
		api.GET("/article/:id", controllers.GetArticlesById)

		api.POST("article/:id/like", controllers.LikeAricleById)
		api.GET("/article/:id/like", controllers.GetLikesById)
	}

	return r
}
