package types

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"

	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/ewkbhex"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GeoPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (g GeoPoint) GormDataType() string {
	return "geometry(Point, 4326)"
}

func (g GeoPoint) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_GeomFromText(?)",
		Vars: []interface{}{fmt.Sprintf("POINT(%f %f)", g.X, g.Y)},
	}
}

func (g *GeoPoint) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	pt, err := ewkbhex.Decode(value.(string))

	if err == nil {
		if p, ok := pt.(*geom.Point); ok {
			g.X = p.X()
			g.Y = p.Y()
		} else {
			return errors.New(fmt.Sprint("Failed to unmarshal geometry:", pt))
		}
	}

	return err
}

func (g GeoPoint) Value() (driver.Value, error) {
	return fmt.Sprintf("POINT(%f %f)", g.X, g.Y), nil
}
