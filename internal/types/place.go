package types

import "gorm.io/gorm"

type Place struct {
	gorm.Model
	Point       GeoPoint  `gorm:"index;not null" json:"point"`
	Name        string    `gorm:"index;not null;size:64" json:"name"`
	Description string    `gorm:"not null;size:1024" json:"description"`
	Avatar      string    `gorm:"not null;size:1024" json:"avatar"`
	Galleries   []Gallery `gorm:"foreignKey:PlaceId;references:ID" json:"galleries"`
}
