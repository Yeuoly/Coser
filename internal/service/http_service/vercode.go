package http_service

import (
	"github.com/Yeuoly/billboards/internal/middleware/auth"
	"github.com/Yeuoly/billboards/internal/types"
)

func EmailVercode(pre_vercode_token string, per_vercode_result string, pre_method string, email string, prompt string, method string) *types.BillboardsResponse {
	// check vercode
	pre_vercode_obj, err := auth.ImageVercodeFromToken(pre_vercode_token)
	if err != nil {
		return types.ErrorResponse(-403, "please complete the image verification")
	}

	if !pre_vercode_obj.Check(pre_method, per_vercode_result) {
		return types.ErrorResponse(-403, "invalid image verification")
	}

	type response struct {
		Token string `json:"token"`
	}

	token, err := auth.SendVercodeEmail(email, prompt, method)
	if err != nil {
		return types.ErrorResponse(-500, err.Error())
	}

	return types.SuccessResponse(response{Token: token})
}

func ImageVercode(method string) *types.BillboardsResponse {
	type response struct {
		Token string `json:"token"`
		Image string `json:"image"`
	}

	image, token, err := auth.GetVercodeImage(method)
	if err != nil {
		return types.ErrorResponse(-500, err.Error())
	}

	return types.SuccessResponse(response{
		Token: token,
		Image: image,
	})
}
