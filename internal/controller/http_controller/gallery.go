package http_controller

import (
	"github.com/Yeuoly/coshub/internal/controller"
	"github.com/gin-gonic/gin"
)

func HandleCreateGallery(c *gin.Context) {
	type request struct {
		PlaceID       uint   `json:"place_id" binding:"required"`
		Name          string `json:"name" binding:"required,max=64"`
		Cosers        string `json:"cosers" binding:"required,max=1024"`
		Photographers string `json:"photographers" binding:"required,max=1024"`
		Description   string `json:"description" binding:"max=1024"`
		Character     string `json:"character" binding:"max=1024"`
		Series        string `json:"series" binding:"max=1024"`
		Tags          []uint `json:"tags"` // tag id
		Key           string `json:"key" binding:"required,max=64"`
	}

	controller.BindRequest(c, func(r request) {
	})
}

func HandleUpdateGallery(c *gin.Context) {
	type request struct {
		ID            uint   `json:"id" binding:"required"`
		PlaceID       uint   `json:"place_id" binding:"required"`
		Name          string `json:"name" binding:"required,max=64"`
		Cosers        string `json:"cosers" binding:"required,max=1024"`
		Photographers string `json:"photographers" binding:"required,max=1024"`
		Description   string `json:"description" binding:"max=1024"`
		Character     string `json:"character" binding:"max=1024"`
		Series        string `json:"series" binding:"max=1024"`
		Tags          []uint `json:"tags"` // tag id
		Key           string `json:"key" binding:"required,max=64"`
	}

	controller.BindRequest(c, func(r request) {
	})
}

func HandleGalleryInfo(c *gin.Context) {
	type request struct {
		ID uint `json:"id" binding:"required" form:"id"`
	}

	controller.BindRequest(c, func(r request) {
	})
}

// search by keyword and tags
func HandleGallerySearch(c *gin.Context) {
	type request struct {
		Keyword string `json:"keyword" form:"keyword"`
	}

	controller.BindRequest(c, func(r request) {
	})
}

// list by place
func HandleGalleryList(c *gin.Context) {
	type request struct {
		PlaceID uint `json:"place_id" form:"place_id"`
	}

	controller.BindRequest(c, func(r request) {
	})
}

// upload images
func HandleGalleryUploadImage(c *gin.Context) {
	type request struct {
		GalleryID uint   `json:"gallery_id" binding:"required"`
		Key       string `json:"key" binding:"required"`
	}

	controller.BindRequest(c, func(r request) {
	})
}

// delete images
func HandleGalleryDeleteImage(c *gin.Context) {
	type request struct {
		GalleryID uint   `json:"gallery_id" binding:"required"`
		Id        uint   `json:"id" binding:"required"`
		Key       string `json:"key" binding:"required"`
	}

	controller.BindRequest(c, func(r request) {
	})
}

// delete gallery
func HandleGalleryDelete(c *gin.Context) {
	type request struct {
		ID  uint   `json:"id" binding:"required"`
		Key string `json:"key" binding:"required"`
	}

	controller.BindRequest(c, func(r request) {
	})
}
