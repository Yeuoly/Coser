package http_service

import (
	"github.com/Yeuoly/billboards/internal/db"
	"github.com/Yeuoly/billboards/internal/types"
)

func AdminSearchUser(user *types.User, username string) *types.BillboardsResponse {
	type response struct {
		Users []types.User `json:"users"`
	}

	users, err := db.GetAll[types.User](
		db.Like("username", username),
		db.OrderBy("id", true),

		db.Page(1, 20),
	)

	if err != nil {
		return types.ErrorResponse(-500, "internal error")
	}

	return types.SuccessResponse(response{
		Users: users,
	})
}

func AdminListUser(user *types.User, page int) *types.BillboardsResponse {
	type response struct {
		Users []types.User `json:"users"`
		Total uint         `json:"total"`
	}

	users, err := db.GetAll[types.User](
		db.Page(page, 20),
		db.OrderBy("id", true),
	)

	if err != nil {
		return types.ErrorResponse(-500, "internal error")
	}

	total, err := db.GetCount[types.User]()
	if err != nil {
		return types.ErrorResponse(-500, "internal error")
	}

	return types.SuccessResponse(response{
		Users: users,
		Total: uint(total),
	})
}
