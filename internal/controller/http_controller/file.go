package http_controller

import (
	"github.com/Yeuoly/coshub/internal/controller"
	"github.com/Yeuoly/coshub/internal/service/http_service"
	"github.com/gin-gonic/gin"
)

func HandleFileUpload(c *gin.Context) {
	type request struct {
		Filename    string `json:"filename" binding:"required"`
		ContentType string `json:"content_type" binding:"required"`
	}

	controller.BindRequest(c, func(r request) {
		c.JSON(200, http_service.UploadFile(r.Filename, r.ContentType, c.ClientIP()))
	})
}

func HandleFileUploadFinish(c *gin.Context) {
	type request struct {
		Url string `json:"url" binding:"required"`
	}

	controller.BindRequest(c, func(r request) {
		c.JSON(200, http_service.FinishUploadFile(r.Url))
	})
}
