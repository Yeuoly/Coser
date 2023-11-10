package router

import (
	"github.com/Yeuoly/coshub/internal/controller/http_controller"
	"github.com/Yeuoly/coshub/internal/middleware"
	"github.com/Yeuoly/coshub/internal/types"
	"github.com/gin-gonic/gin"
)

func Setup(eng *gin.Engine, config *types.CoshubGlobalConfigurations) {
	eng.Use(middleware.Cors())

	eng.POST("/v1/file/upload", http_controller.HandleFileUpload)
	eng.POST("/v1/file/upload/finish", http_controller.HandleFileUploadFinish)

	eng.POST("/v1/place/create", http_controller.HandlePlaceCreate)
	eng.POST("/v1/place/update", http_controller.HandlePlaceUpdate)
	eng.GET("/v1/place/info", http_controller.HandlePlaceInfo)
	eng.GET("/v1/place/list", http_controller.HandlePlaceList)
	eng.GET("/v1/place/nearby", http_controller.HandlePlaceNearby)

	eng.POST("/v1/tag/create", http_controller.HandleCreateTag)
	eng.GET("/v1/tag/search", http_controller.HandleSearchTag)

	eng.POST("/v1/gallery/create", http_controller.HandleCreateGallery)
	eng.POST("/v1/gallery/update", http_controller.HandleUpdateGallery)
	eng.GET("/v1/gallery/info", http_controller.HandleGalleryInfo)
	eng.GET("/v1/gallery/search", http_controller.HandleGallerySearch)
	eng.POST("/v1/gallery/upload", http_controller.HandleGalleryUploadImage)
	eng.POST("/v1/gallery/delete", http_controller.HandleGalleryDeleteImage)
}
