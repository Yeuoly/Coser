package db

import (
	"fmt"

	"github.com/Yeuoly/coshub/internal/types"
	"gorm.io/gorm"
)

/*
GeoDistance is a function that returns a gorm.DB filter that filters the distance between the center and the field to be less than or equal to the distance.
distance is in meters and center is a GeoPoint.
You should make sure that the field is a GeoPoint.
*/
func GeoDistance(field string, center *types.GeoPoint, distance float64) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("ST_DistanceSphere(%s, ?) <= ?", field), center, distance)
	}
}
