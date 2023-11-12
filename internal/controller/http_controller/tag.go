package http_controller

import (
	"github.com/Yeuoly/coshub/internal/controller"
	"github.com/Yeuoly/coshub/internal/service/http_service"
	"github.com/gin-gonic/gin"
)

func HandleCreateTag(c *gin.Context) {
	type request struct {
		Name string `json:"name" binding:"required,max=64"`
	}

	controller.BindRequest(c, func(r request) {
		c.JSON(200, http_service.CreateTag(r.Name))
	})
}

func HandleSearchTag(c *gin.Context) {
	type request struct {
		Keyword string `json:"keyword" form:"keyword" binding:"required,max=64"`
	}

	controller.BindRequest(c, func(r request) {
		c.JSON(200, http_service.SearchTag(r.Keyword))
	})
}
