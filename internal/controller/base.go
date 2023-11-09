package controller

import (
	"github.com/Yeuoly/coshub/internal/middleware"
	"github.com/Yeuoly/coshub/internal/types"
	"github.com/gin-gonic/gin"
)

func BindRequest[T any](r *gin.Context, success func(T)) {
	var request T
	var err error

	context_type := r.GetHeader("Content-Type")
	if context_type == "application/json" {
		err = r.BindJSON(&request)
	} else {
		err = r.ShouldBind(&request)
	}

	if err != nil {
		resp := types.ErrorResponse(-400, err.Error())
		r.JSON(200, resp)
		return
	}
	success(request)
}

func AfterAuthorization[T any](r *gin.Context, success func(user *types.User, request T), require_admin ...bool) {
	user_interface, exists := r.Get(middleware.CONTEXT_KEY_USER)
	if !exists {
		resp := types.ErrorResponse(-401, "Unauthorized")
		r.JSON(200, resp)
		return
	}

	user, ok := user_interface.(*types.User)
	if !ok {
		resp := types.ErrorResponse(-401, "Unauthorized")
		r.JSON(200, resp)
		return
	}

	if len(require_admin) > 0 && require_admin[0] {
		if !user.IsAdmin() {
			resp := types.ErrorResponse(-401, "Unauthorized")
			r.JSON(200, resp)
			return
		}
	}

	BindRequest[T](r, func(request T) {
		success(user, request)
	})
}
