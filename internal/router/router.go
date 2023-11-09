package router

import (
	"github.com/Yeuoly/coshub/internal/middleware"
	"github.com/Yeuoly/coshub/internal/types"
	"github.com/gin-gonic/gin"
)

func Setup(eng *gin.Engine, config *types.CoshubGlobalConfigurations) {
	eng.Use(middleware.Cors())
	eng.Use(middleware.Auth())

	// eng.POST(static.ROUTE_USER_LOGIN, http_controller.HandleUserLogin)
	// eng.POST(static.ROUTE_USER_LOGIN_GITHUB, http_controller.HandleUserLoginWithGithub)
	// eng.POST(static.ROUTE_USER_LOGIN_KASUMI_ACCEPT, http_controller.HandleUserRequestLoginKasumiAccept)
	// eng.POST(static.ROUTE_USER_LOGIN_KASUMI_CHECK, http_controller.HandleUserRequestLoginKasumiCheck)
	// eng.POST(static.ROUTE_USER_REG, http_controller.HandleUserRegister)
	// eng.GET(static.ROUTE_USER_CHECK, http_controller.HandleUserCheck)
	// eng.POST(static.ROUTE_USER_UPDATE, http_controller.HandleUserUpdate)
	// eng.GET(static.ROUTE_USER_PROFILE, http_controller.HandleUserProfile)
	// eng.GET(static.ROUTE_USER_ADMIN_SEARCH, http_controller.HandleAdminSearchUser)
	// eng.GET(static.ROUTE_USER_ADMIN_LIST, http_controller.HandleAdminListUser)

	// eng.GET(static.ROUTE_VERCODE_LOGIN, http_controller.HandleEmailLoginVercode)
	// eng.GET(static.ROUTE_VERCODE_REG_PRE, http_controller.HandleEmailRegisterPreVercode)
	// eng.POST(static.ROUTE_VERCODE_REG, http_controller.HandleEmailRegisterVercode)
}
