package types

type Tag struct {
	Model
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"index;not null;size:64" json:"name"`
	Galleries []Gallery `gorm:"many2many:gallery_tags_relations" json:"galleries"`
}
