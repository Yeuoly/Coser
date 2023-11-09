package http_controller

import (
	"github.com/Yeuoly/coshub/internal/controller"
	"github.com/gin-gonic/gin"
)

func HandlePlaceCreate(c *gin.Context) {
	type request struct {
		Name        string  `json:"name" binding:"required,max=64"`
		Description string  `json:"description" binding:"max=1024"`
		Lat         float64 `json:"lat" binding:"required"`
		Lng         float64 `json:"lng" binding:"required"`
		Avatar      string  `json:"avatar" binding:"max=1024"`
		Key         string  `json:"key" binding:"required,max=64"`
	}

	controller.BindRequest(c, func(r request) {
	})
}

func HandlePlaceUpdate(c *gin.Context) {
	type request struct {
		Name        string  `json:"name" binding:"required,max=64"`
		Description string  `json:"description" binding:"max=1024"`
		Lat         float64 `json:"lat" binding:"required"`
		Lng         float64 `json:"lng" binding:"required"`
		Avatar      string  `json:"avatar" binding:"max=1024"`
		Key         string  `json:"key" binding:"required,max=64"`
	}

	controller.BindRequest(c, func(r request) {
	})
}

func HandlePlaceInfo(c *gin.Context) {
	type request struct {
		ID uint `json:"id" binding:"required"`
	}

	controller.BindRequest(c, func(r request) {
	})
}

func HandlePlaceList(c *gin.Context) {
	type request struct {
		Keyword string `json:"keyword"`
	}

	controller.BindRequest(c, func(r request) {
	})
}

func HandlePlaceNearby(c *gin.Context) {
	type request struct {
		Lat      float64 `json:"lat" binding:"required"`
		Lng      float64 `json:"lng" binding:"required"`
		Distance float64 `json:"distance" binding:"required"` // in meters
	}

	controller.BindRequest(c, func(r request) {
	})
}
