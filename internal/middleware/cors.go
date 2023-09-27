package middleware

import (
	"strings"

	"github.com/Yeuoly/billboards/internal/static"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	config := static.GetBillboardsGlobalConfigurations()
	return func(c *gin.Context) {
		// get origin
		origin := c.Request.Header.Get("Origin")
		// check if origin is allowed
		for _, v := range config.Cors.AllowOrigins {
			if v == origin {
				c.Header("Access-Control-Allow-Origin", origin)
				c.Header("Access-Control-Allow-Methods", strings.Join(config.Cors.AllowMethods, ","))
				c.Header("Access-Control-Allow-Headers", strings.Join(config.Cors.AllowHeaders, ","))
				c.Header("Access-Control-Allow-Credentials", "true")
				c.Header("Access-Control-Max-Age", "86400")
				// if options, return
				if c.Request.Method == "OPTIONS" {
					c.AbortWithStatus(200)
				}
				return
			}
		}
	}
}
