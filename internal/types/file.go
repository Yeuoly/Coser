package types

type File struct {
	Model
	Url string `gorm:"index;not null;size:1024" json:"url"`
	Ip  string `gorm:"not null;size:64" json:"ip"` // ip address of the uploader, for logging
}
