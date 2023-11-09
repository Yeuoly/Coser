package http_controller

import (
	"github.com/Yeuoly/coshub/internal/controller"
	service "github.com/Yeuoly/coshub/internal/service/http_service"
	"github.com/Yeuoly/coshub/internal/types"
	"github.com/gin-gonic/gin"
)

func HandleUserCheck(c *gin.Context) {
	controller.AfterAuthorization[struct{}](c, func(user *types.User, request struct{}) {
		c.JSON(200, service.UserCheck(user))
	})
}

func HandleUserLogin(c *gin.Context) {
	type request struct {
		Email        string `json:"email" binding:"required,email"`
		Password     string `json:"password" binding:"required,len=32"`
		VercodeToken string `json:"vercode_token" binding:"required,lte=512"`
		Vercode      string `json:"vercode" binding:"required,lte=32"`
	}
	controller.BindRequest[request](c, func(req request) {
		c.JSON(200, service.UserLogin(req.Email, req.Password, req.VercodeToken, req.Vercode))
	})
}

func HandleUserRegister(c *gin.Context) {
	type request struct {
		Username     string `json:"username" binding:"required,lte=32"`
		Password     string `json:"password" binding:"required,len=32"`
		VercodeToken string `json:"vercode_token" binding:"required,lte=512"`
		Vercode      string `json:"vercode" binding:"required,lte=32"`
		InviteCode   string `json:"invite_code"`
	}
	controller.BindRequest[request](c, func(req request) {
		c.JSON(200, service.UserRegister(req.Username, req.Password, req.Vercode, req.VercodeToken, req.InviteCode))
	})
}

func HandleUserLoginWithGithub(c *gin.Context) {
	type request struct {
		Code string `json:"code" binding:"required,lte=128"`
	}

	controller.BindRequest[request](c, func(req request) {
		c.JSON(200, service.UserLoginGithub(req.Code))
	})
}

func HandleUserProfile(c *gin.Context) {
	controller.AfterAuthorization[struct{}](
		c, func(user *types.User, request struct{}) {
			c.JSON(200, service.UserProfile(user))
		},
	)
}

func HandleUserUpdate(c *gin.Context) {
	type request struct {
		Username string `json:"username" binding:"required,lte=32"`
		Avatar   string `json:"avatar" binding:"lte=128"`
		Sign     string `json:"sign" binding:"lte=128"`
	}

	controller.AfterAuthorization[request](c, func(user *types.User, req request) {
		c.JSON(200, service.UserUpdate(user, req.Username, req.Avatar, req.Sign))
	})
}

func HandleUserRequestLoginKasumiAccept(c *gin.Context) {
	type request struct {
		Token string `json:"token" binding:"required,len=32"`
	}

	controller.AfterAuthorization[request](c, func(user *types.User, req request) {
		c.JSON(200, service.UserRequestLoginKasumiAccept(user, req.Token))
	})
}

func HandleUserRequestLoginKasumiCheck(c *gin.Context) {
	type request struct {
		Token string `json:"token" binding:"required,len=32"`
	}

	controller.BindRequest[request](c, func(req request) {
		c.JSON(200, service.UserRequestLoginKasumiCheck(req.Token))
	})
}
