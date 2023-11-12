package http_service

import (
	"github.com/Yeuoly/coshub/internal/db"
	"github.com/Yeuoly/coshub/internal/types"
)

func CreateTag(name string) *types.CoshubResponse {
	type response struct {
		ID uint `json:"id"`
	}

	tag, err := db.GetOne[types.Tag](
		db.Equal("name", name),
	)

	if err != nil && err != db.ErrDatabaseNotFound {
		return types.ErrorResponse(-500, "internal error")
	}

	if err == nil {
		return types.SuccessResponse(response{
			ID: tag.ID,
		})
	}

	tag = types.Tag{
		Name: name,
	}

	err = db.Create(&tag)

	if err != nil {
		return types.ErrorResponse(-500, "internal error")
	}

	return types.SuccessResponse(response{
		ID: tag.ID,
	})
}

func SearchTag(keyword string) *types.CoshubResponse {
	type response struct {
		Tags []types.Tag `json:"tags"`
	}

	tags, err := db.GetAll[types.Tag](
		db.Like("name", keyword),
		db.Page(1, 10),
	)

	if err != nil {
		return types.ErrorResponse(-500, "internal error")
	}

	return types.SuccessResponse(response{
		Tags: tags,
	})
}
