package http_service

import (
	"github.com/Yeuoly/coshub/internal/db"
	"github.com/Yeuoly/coshub/internal/types"
	"github.com/Yeuoly/coshub/internal/utils/coordinate"
)

func CreatePlace(name string, description string, lat float64, lng float64, avatar string, key string) *types.CoshubResponse {
	type response struct {
		ID uint `json:"id"`
	}

	// check if already exists
	place, err := db.GetOne[types.Place](
		db.Equal("name", name),
	)

	if err != nil && err != db.ErrDatabaseNotFound {
		return types.ErrorResponse(-500, "internal error")
	}

	if err == nil {
		return types.ErrorResponse(-400, "有人已经占用了这个地点的名字")
	}

	x, y := coordinate.LL2XY(lng, lat)

	// create
	place = types.Place{
		Name:        name,
		Description: description,
		Point: types.GeoPoint{
			X: x,
			Y: y,
		},
		Avatar: avatar,
		Key:    key,
	}

	err = db.Create(&place)

	if err != nil {
		return types.ErrorResponse(-500, "internal error")
	}

	return types.SuccessResponse(response{
		ID: place.ID,
	})
}

func GetNearbyPlaces(lat float64, lng float64) *types.CoshubResponse {
	type response struct {
		Places []types.Place `json:"places"`
	}

	x, y := coordinate.LL2XY(lng, lat)

	places, err := db.GetAll[types.Place](
		db.GeoDistance("point", &types.GeoPoint{
			X: x,
			Y: y,
		}, 100000),
	)

	if err != nil {
		return types.ErrorResponse(-500, "internal error")
	}

	for i := range places {
		places[i].ClearSensitive()
	}

	return types.SuccessResponse(response{
		Places: places,
	})
}
