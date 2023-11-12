package types

type Tag struct {
	Model
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"index;not null;size:64" json:"name"`
}
