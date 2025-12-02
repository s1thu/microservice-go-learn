package router

import (
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
	v1 := r.Group("/api/v1")
	{
		albums := v1.Group("/albums")
		{
			albums.GET("/", service.GetAllAlbums)
			albums.POST("/", service.PostAlbum)
			albums.GET("/:id", service.GetAllAlbumByID)
			albums.PUT("/:id", service.UpdateAlbum)
			albums.DELETE("/:id", service.DeleteAlbum)
			albums.PATCH("/:id", service.PatchAlbum)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
