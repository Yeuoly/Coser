/**
 * 
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
	Ip            string       `gorm:"not null;size:64" json:"ip"`  // ip address of the uploader, for logging
	Key           string       `gorm:"not null;size:64" json:"key"` // key of the gallery, only the browser with the key can change the gallery
}


type Image struct {
	gorm.Model
	Url         string `gorm:"not null;size:1024" json:"url"`
	GalleryId   uint   `gorm:"index;not null" json:"gallery_id"`
	Camera      string `gorm:"not null;size:64" json:"camera"`
	Lens        string `gorm:"not null;size:64" json:"lens"`
	FocalLength string `gorm:"not null;size:64" json:"focal_length"`
	Ip          string `gorm:"not null;size:64" json:"ip"` // ip address of the uploader, for logging
}

 */

export type Image = {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string;
    url: string;
    camera: string;
    lens: string;
    focal_length: string;
}

export type GalleryTag = {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string;
    name: string;
}

export type Gallery = {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string;
    name: string;
    cosers: string;
    photographers: string;
    description: string;
    character: string;
    series: string;
    place_id: number;
    place: Place;
    tags: GalleryTag[];
    images: Image[];
    key: string;
}

export type Place = {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string;
    point: {
        x: number;
        y: number;
    };
    name: string;
    description: string;
    avatar: string;
    galleries?: Gallery[];
}