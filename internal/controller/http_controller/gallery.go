package http_controller

import (
	"github.com/Yeuoly/coshub/internal/controller"
	"github.com/Yeuoly/coshub/internal/service/http_service"
	"github.com/gin-gonic/gin"
)

func HandleCreateGallery(c *gin.Context) {
	type request struct {
		PlaceID       uint   `json:"place_id" binding:"required"`
		Name          string `json:"name" binding:"required,max=64"`
		Cosers        string `json:"cosers" binding:"max=1024"`
		Photographers string `json:"photographers" binding:"max=1024"`
		Description   string `json:"description" binding:"max=1024"`
		Character     string `json:"character" binding:"max=1024"`
		Series        string `json:"series" binding:"max=1024"`
		Tags          []uint `json:"tags"` // tag id
		Key           string `json:"key" binding:"required,max=64"`
	}

	controller.BindRequest(c, func(r request) {
		c.JSON(200, http_service.CreateGallery(r.PlaceID, r.Name, r.Cosers, r.Photographers, r.Description, r.Character, r.Series, r.Tags, r.Key, c.ClientIP()))
	})
}

func HandleUpdateGallery(c *gin.Context) {
	type request struct {
		ID            uint   `json:"id" binding:"required"`
		PlaceID       uint   `json:"place_id" binding:"required"`
		Name          string `json:"name" binding:"required,max=64"`
		Cosers        string `json:"cosers" binding:"max=1024"`
		Photographers string `json:"photographers" binding:"max=1024"`
		Description   string `json:"description" binding:"max=1024"`
		Character     string `json:"character" binding:"max=1024"`
		Series        string `json:"series" binding:"max=1024"`
		Tags          []uint `json:"tags"` // tag id
		Key           string `json:"key" binding:"required,max=64"`
	}

	controller.BindRequest(c, func(r request) {
		c.JSON(200, http_service.UpdateGallery(r.ID, r.PlaceID, r.Name, r.Cosers, r.Photographers, r.Description, r.Character, r.Series, r.Tags, r.Key, c.ClientIP()))
	})
}

func HandleGalleryInfo(c *gin.Context) {
	type request struct {
		ID uint `json:"id" binding:"required" form:"id"`
	}

	controller.BindRequest(c, func(r request) {
		c.JSON(200, http_service.GetGalleryInfo(r.ID))
	})
}

// search by keyword and tags
func HandleGallerySearch(c *gin.Context) {
	type request struct {
		Keyword string `json:"keyword" form:"keyword"`
	}

	controller.BindRequest(c, func(r request) {
		c.JSON(200, http_service.SearchGallery(r.Keyword))
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
		GalleryID      uint   `json:"gallery_id" binding:"required"`
		Filename       string `json:"filename" binding:"required"`
		Key            string `json:"key" binding:"required"`
		ContenType     string `json:"content_type" binding:"required"`
		Camera         string `json:"camera" binding:"max=64"`
		Lens           string `json:"lens" binding:"max=64"`
		AperatureValue string `json:"aperture_value" binding:"max=64"`
		ExposureTime   string `json:"exposure_time" binding:"max=64"`
		ISO            string `json:"iso" binding:"max=64"`
		FocalLength    string `json:"focal_length" binding:"max=64"`
		Width          uint   `json:"width" binding:""`
		Height         uint   `json:"height" binding:""`
	}

	controller.BindRequest(c, func(r request) {
		c.JSON(200, http_service.UploadGallery(
			r.GalleryID, r.Filename, r.ContenType,
			r.Camera, r.Lens, r.FocalLength, r.AperatureValue, r.ExposureTime, r.ISO,
			r.Width, r.Height,
			c.ClientIP(),
		))
	})
}

// update image info
func HandleGalleryUpdateImage(c *gin.Context) {
	type request struct {
		GalleryID    uint   `json:"gallery_id" binding:"required"`
		ID           uint   `json:"id" binding:"required"`
		Camera       string `json:"camera" binding:"max=64"`
		Lens         string `json:"lens" binding:"max=64"`
		Aperature    string `json:"aperture_value" binding:"max=64"`
		ExposureTime string `json:"exposure_time" binding:"max=64"`
		ISO          string `json:"iso" binding:"max=64"`
		FocalLength  string `json:"focal_length" binding:"max=64"`
		Key          string `json:"key" binding:"required"`
	}

	controller.BindRequest(c, func(r request) {
		c.JSON(200, http_service.UpdateGalleryImage(
			r.ID, r.GalleryID,
			r.Camera, r.Lens, r.FocalLength, r.Aperature, r.ExposureTime, r.ISO,
			r.Key,
		))
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
		c.JSON(200, http_service.DeleteGalleryImage(r.GalleryID, r.Id, r.Key))
	})
}

// delete gallery
func HandleGalleryDelete(c *gin.Context) {
	type request struct {
		ID  uint   `json:"id" binding:"required"`
		Key string `json:"key" binding:"required"`
	}

	controller.BindRequest(c, func(r request) {
		c.JSON(200, http_service.DeleteGallery(r.ID, r.Key))
	})
}
