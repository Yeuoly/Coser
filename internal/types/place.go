package types

import "gorm.io/gorm"

type Place struct {
	gorm.Model
	Point       GeoPoint  `gorm:"index;not null" json:"point"`
	Name        string    `gorm:"index;not null;size:64" json:"name"`
	Description string    `gorm:"not null;size:1024" json:"description"`
	Avatar      string    `gorm:"not null;size:1024" json:"avatar"`
	Galleries   []Gallery `gorm:"foreignKey:PlaceId;references:ID" json:"galleries"`
	Key         string    `gorm:"not null;size:64" json:"key"` // key of the place, only the browser with the key can change the place
	Ip          string    `gorm:"not null;size:64" json:"ip"`  // ip address of the uploader, for logging
}

func (p *Place) ClearSensitive() {
	p.Key = ""
	p.Ip = ""
	for i := range p.Galleries {
		p.Galleries[i].ClearSensitive()
	}
}
