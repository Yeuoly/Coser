package types

import "gorm.io/gorm"

type GalleryTag struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"index;not null;size:64" json:"name"`
}
