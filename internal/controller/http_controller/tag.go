package http_controller

import (
	"github.com/Yeuoly/coshub/internal/controller"
	"github.com/gin-gonic/gin"
)

func HandleCreateTag(c *gin.Context) {
	type request struct {
		Name string `json:"name" binding:"required,max=64"`
	}

	controller.BindRequest(c, func(r request) {
	})
}

func HandleSearchTag(c *gin.Context) {
	type request struct {
		Keyword string `json:"keyword" form:"keyword" binding:"required,max=64"`
	}

	controller.BindRequest(c, func(r request) {
	})
}
