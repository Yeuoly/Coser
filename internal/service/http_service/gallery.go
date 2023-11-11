package http_service

import (
	"strings"
	"time"

	"github.com/Yeuoly/coshub/internal/db"
	"github.com/Yeuoly/coshub/internal/types"
	"github.com/Yeuoly/coshub/internal/utils/oss"
	"github.com/Yeuoly/coshub/internal/utils/slice"
	"gorm.io/gorm"
)

func CreateGallery(placeID uint, name, cosers, photographers, description, character, series string, tags []uint, key string, ip string) *types.CoshubResponse {
	type response struct {
		ID uint `json:"id"`
	}

	// check if every tags exists
	tags_obj, err := db.GetAll[types.Tag](
		db.InArray("id", slice.Map(func(r uint) interface{} {
			return r
		}, tags)),
	)

	if err != nil {
		return types.ErrorResponse(-500, "internal error")
	}

	if len(tags_obj) != len(tags) {
		return types.ErrorResponse(-400, "有不存在的标签")
	}

	// check if place exists
	_, err = db.GetOne[types.Place](
		db.Equal("id", placeID),
	)

	if err != nil {
		return types.ErrorResponse(-500, "internal error")
	}

	gallery := &types.Gallery{
		PlaceId:       placeID,
		Name:          name,
		Cosers:        cosers,
		Photographers: photographers,
		Description:   description,
		Character:     character,
		Series:        series,
		Ip:            ip,
		Key:           key,
	}

	err = db.WithTransaction(func(tx *gorm.DB) error {
		err = db.Create(gallery, tx)

		if err != nil {
			return err
		}

		// append tags association
		for _, tag := range tags_obj {
			err = db.AppendAssociation(gallery, "Tags", tag, tx)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return types.ErrorResponse(-500, "internal error")
	}

	return types.SuccessResponse(response{gallery.ID})
}

func GetGalleryInfo(id uint) *types.CoshubResponse {
	type response struct {
		Gallery types.Gallery `json:"gallery"`
	}

	gallery, err := db.GetOne[types.Gallery](
		db.Equal("id", id),
		db.Preload("Images"),
		db.Preload("Tags"),
	)

	if err != nil {
		return types.ErrorResponse(-500, "internal error")
	}

	return types.SuccessResponse(response{
		Gallery: gallery,
	})
}

func UploadGallery(galleryID uint, filename string, contentType string, camera string, lens string, focalLength string, ip string) *types.CoshubResponse {
	type response struct {
		Url string `json:"url"`
		Id  uint   `json:"id"`
	}

	gallery, err := db.GetOne[types.Gallery](
		db.Equal("id", galleryID),
	)

	if err != nil {
		return types.ErrorResponse(-500, "internal error")
	}

	url, err := oss.RequestTempraryWriteUrl(filename, time.Minute*5, contentType, oss.OssAutoDatePath("md5"))

	if err != nil {
		return types.ErrorResponse(-500, "internal error")
	}

	file := &types.File{
		Url: url,
		Ip:  ip,
	}

	err = db.Create(file)
	if err != nil {
		return types.ErrorResponse(-500, "internal error")
	}

	storage_url := strings.Split(url, "?")[0]
	image := &types.Image{
		Url:         storage_url,
		GalleryId:   gallery.ID,
		Camera:      camera,
		Lens:        lens,
		FocalLength: focalLength,
		Ip:          ip,
	}

	err = db.Create(image)
	if err != nil {
		return types.ErrorResponse(-500, "internal error")
	}

	return types.SuccessResponse(response{
		Url: url,
		Id:  image.ID,
	})
}

func DeleteGallery(id uint, key string) *types.CoshubResponse {
	gallery, err := db.GetOne[types.Gallery](
		db.Equal("id", id),
	)

	if err != nil {
		return types.ErrorResponse(-500, "internal error")
	}

	if gallery.Key != key {
		return types.ErrorResponse(-403, "key错误")
	}

	err = db.Delete(&gallery)

	if err != nil {
		return types.ErrorResponse(-500, "internal error")
	}

	return types.SuccessResponse(nil)
}
