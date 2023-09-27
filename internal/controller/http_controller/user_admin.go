package http_controller

import (
	"github.com/Yeuoly/billboards/internal/controller"
	"github.com/Yeuoly/billboards/internal/service/http_service"
	"github.com/Yeuoly/billboards/internal/types"
	"github.com/gin-gonic/gin"
)

func HandleAdminSearchUser(c *gin.Context) {
	type request struct {
		Username string `json:"username" form:"username" binding:"required,lte=32"`
	}

	controller.AfterAuthorization[request](c, func(user *types.User, req request) {
		c.JSON(200, http_service.AdminSearchUser(user, req.Username))
	}, true)
}

func HandleAdminListUser(c *gin.Context) {
	type request struct {
		Page int `json:"page" form:"page" binding:"required"`
	}

	controller.AfterAuthorization[request](c, func(user *types.User, req request) {
		c.JSON(200, http_service.AdminListUser(user, req.Page))
	}, true)
}
