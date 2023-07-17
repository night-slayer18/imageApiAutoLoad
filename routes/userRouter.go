package routes

import (
	controller "Uploader/controllers"
	"Uploader/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.POST("/upload/:image_name", controller.UploadFile())
	incomingRoutes.GET("/download/:image_name", controller.DownloadFile())
	incomingRoutes.GET("/latestpost", controller.DFile())
}
