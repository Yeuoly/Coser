package http_controller

import (
	"fmt"

	"github.com/Yeuoly/billboards/internal/controller"
	"github.com/Yeuoly/billboards/internal/service/http_service"
	"github.com/gin-gonic/gin"
)

func HandleEmailLoginVercode(c *gin.Context) {
	type reqeust struct{}
	controller.BindRequest[reqeust](c, func(r reqeust) {
		c.JSON(200, http_service.ImageVercode("normal_login"))
	})
}

func HandleEmailRegisterPreVercode(c *gin.Context) {
	type reqeust struct{}
	controller.BindRequest[reqeust](c, func(r reqeust) {
		c.JSON(200, http_service.ImageVercode("normal_register"))
	})
}

func HandleEmailRegisterVercode(c *gin.Context) {
	type reqeust struct {
		PreVercodeToken  string `json:"pre_vercode_token" binding:"required"`
		PreVercodeResult string `json:"pre_vercode_result" binding:"required"`
		Email            string `json:"email" binding:"required"`
	}
	controller.BindRequest[reqeust](c, func(r reqeust) {
		c.JSON(200, http_service.EmailVercode(r.PreVercodeToken, r.PreVercodeResult, "normal_register", r.Email, fmt.Sprintf(
			"You are trying to register a miduoduo account with your email %s, please enter the verification code below to continue."+
				"Remember, do not tell anyone this code, including the staff of miduoduo, otherwise your account may be stolen.",
			r.Email,
		), "normal_register"))
	})
}
