package types

import "gorm.io/gorm"

type Gallery struct {
	gorm.Model
	Name          string       `gorm:"index;not null;size:64" json:"name"`
	Cosers        string       `gorm:"not null;size:1024" json:"cosers"`        // CN
	Photographers string       `gorm:"not null;size:1024" json:"photographers"` // CN
	Description   string       `gorm:"not null;size:1024" json:"description"`
	Character     string       `gorm:"not null;size:1024" json:"character"` // name of the character
	Series        string       `gorm:"not null;size:1024" json:"series"`    // name of the series
	PlaceId       uint         `gorm:"index;not null" json:"place_id"`
	Place         Place        `gorm:"foreignKey:PlaceId;references:ID" json:"place"`
	Tags          []GalleryTag `gorm:"many2many:gallery_tags;" json:"tags"`
	Images        []Image      `gorm:"foreignKey:GalleryId;references:ID" json:"images"`
}