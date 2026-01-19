package router

import (
	"example/go-web-gin/config"
	"example/go-web-gin/database"
	"example/go-web-gin/handler"
	"example/go-web-gin/middleware"
	"example/go-web-gin/repositories"
	"example/go-web-gin/service"

	"example/go-web-gin/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterRoutes registers all application routes on the provided engine.
func RegisterRoutes(r *gin.Engine) {

	// configure swagger info so UI calls the correct server and base path
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http"}

	database.ConnectDB()
	repo := repositories.NewAlbumRepoImpl(database.DB)
	albumService := service.NewAlbumService(repo)
	albumHandler := handler.NewAlbumHandler(albumService)

	authService := service.NewAuthService([]byte(config.AppConfig.JWTSecret))
	authHandler := handler.NewAuthHandler(authService)
	v1 := r.Group("/api/v1")
	{
		// 🔓 Public routes
		v1.POST("/auth/login", authHandler.Login)

		// 🔒 Protected routes
		protected := v1.Group("/hello")
		protected.Use(middleware.JWTAuth(authService))
		{
			protected.GET("/protected", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "You have accessed a protected route"})
			})
		}
		albums := v1.Group("/albums")
		{
			albums.GET("/", albumHandler.GetAllAlbums)
			albums.POST("/", albumHandler.PostAlbum)
			albums.GET("/:id", albumHandler.GetAlbumByID)
			albums.PUT("/:id", albumHandler.UpdateAlbum)
			albums.DELETE("/:id", albumHandler.DeleteAlbum)
			albums.PATCH("/:id", albumHandler.PatchAlbum)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
