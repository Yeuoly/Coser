package types

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Url string `gorm:"index;not null;size:1024" json:"url"`
	Ip  string `gorm:"not null;size:64" json:"ip"` // ip address of the uploader, for logging
}
