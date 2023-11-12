package http_controller

import (
	"github.com/Yeuoly/coshub/internal/controller"
	"github.com/Yeuoly/coshub/internal/service/http_service"
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
		c.JSON(200, http_service.CreatePlace(r.Name, r.Description, r.Lat, r.Lng, r.Avatar, r.Key))
	})
}

func HandlePlaceUpdate(c *gin.Context) {
	type request struct {
		ID          uint    `json:"id" binding:"required"`
		Name        string  `json:"name" binding:"required,max=64"`
		Description string  `json:"description" binding:"max=1024"`
		Lat         float64 `json:"lat" binding:"required"`
		Lng         float64 `json:"lng" binding:"required"`
		Avatar      string  `json:"avatar" binding:"max=1024"`
		Key         string  `json:"key" binding:"required,max=64"`
	}

	controller.BindRequest(c, func(r request) {
		c.JSON(200, http_service.UpdatePlace(r.ID, r.Name, r.Description, r.Lat, r.Lng, r.Avatar, r.Key))
	})
}

func HandlePlaceInfo(c *gin.Context) {
	type request struct {
		ID uint `json:"id" binding:"required" form:"id"`
	}

	controller.BindRequest(c, func(r request) {
		c.JSON(200, http_service.GetPlaceInfo(r.ID))
	})
}

func HandlePlaceDelete(c *gin.Context) {
	type request struct {
		ID  uint   `json:"id" binding:"required" form:"id"`
		Key string `json:"key" binding:"required,max=64"`
	}

	controller.BindRequest(c, func(r request) {
		c.JSON(200, http_service.DeletePlace(r.ID, r.Key))
	})
}

func HandlePlaceList(c *gin.Context) {
	type request struct {
		Keyword string `json:"keyword" form:"keyword"`
	}

	controller.BindRequest(c, func(r request) {
	})
}

func HandlePlaceNearby(c *gin.Context) {
	type request struct {
		Lat float64 `json:"lat" binding:"required" form:"lat"`
		Lng float64 `json:"lng" binding:"required" form:"lng"`
	}

	controller.BindRequest(c, func(r request) {
		c.JSON(200, http_service.GetNearbyPlaces(r.Lat, r.Lng))
	})
}
