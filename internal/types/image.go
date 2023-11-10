package types

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Url         string `gorm:"not null;size:1024" json:"url"`
	GalleryId   uint   `gorm:"index;not null" json:"gallery_id"`
	Camera      string `gorm:"not null;size:64" json:"camera"`
	Lens        string `gorm:"not null;size:64" json:"lens"`
	FocalLength string `gorm:"not null;size:64" json:"focal_length"`
	Ip          string `gorm:"not null;size:64" json:"ip"` // ip address of the uploader, for logging
}

func (i *Image) ClearSensitive() {
	i.Ip = ""
}
