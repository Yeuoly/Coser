package types

type Image struct {
	Model
	Url          string `gorm:"not null;size:1024" json:"url"`
	GalleryId    uint   `gorm:"index;not null" json:"gallery_id"`
	Camera       string `gorm:"not null;size:64" json:"camera"`
	Lens         string `gorm:"not null;size:64" json:"lens"`
	FocalLength  string `gorm:"not null;size:64" json:"focal_length"`
	Aperature    string `gorm:"not null;size:64;default:''" json:"aperature"`
	ExposureTime string `gorm:"not null;size:64;default:''" json:"exposure_time"`
	ISO          string `gorm:"not null;size:64;default:''" json:"iso"`
	Width        uint   `gorm:"not null;default:0" json:"width"`
	Height       uint   `gorm:"not null;default:0" json:"height"`
	Ip           string `gorm:"not null;size:64" json:"ip"` // ip address of the uploader, for logging
}

func (i *Image) ClearSensitive() {
	i.Ip = ""
}
