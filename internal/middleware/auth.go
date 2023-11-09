package middleware

import (
	"strconv"

	"github.com/Yeuoly/coshub/internal/db/model"
	"github.com/Yeuoly/coshub/internal/middleware/auth"
	"github.com/gin-gonic/gin"
)

const (
	CONTEXT_KEY_USER = "user"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid_string := c.GetHeader("Billboards-User-Id")
		uid, err := strconv.Atoi(uid_string)
		if err != nil {
			return
		}

		authoration := c.GetHeader("Billboards-User-Token")
		if authoration == "" {
			return
		}

		tokens := auth.GetUserAuthTokens(uint(uid))
		if tokens == nil {
			return
		}

		success, _ := auth.CheckLoginToken(uint(uid), authoration)
		if success {
			user, err := model.GetUser(uint(uid))
			if err != nil {
				return
			}
			c.Set(CONTEXT_KEY_USER, user)
		}
	}
}
