package http_service

import (
	"time"

	"github.com/Yeuoly/coshub/internal/db"
	"github.com/Yeuoly/coshub/internal/db/model"
	"github.com/Yeuoly/coshub/internal/middleware/auth"
	"github.com/Yeuoly/coshub/internal/middleware/user"
	"github.com/Yeuoly/coshub/internal/static"
	"github.com/Yeuoly/coshub/internal/types"
	"github.com/Yeuoly/coshub/internal/utils/math"
)

func UserCheck(user *types.User) *types.CoshubResponse {
	var response struct {
		Username string `json:"username"`
		Role     string `json:"role"`
		Uid      uint   `json:"uid"`
	}
	response.Username = user.ToResponse().Username
	response.Role = user.ToResponse().Role
	response.Uid = user.ID

	return types.SuccessResponse(response)
}

func UserLogin(email, password, vercode_token, vercode string) *types.CoshubResponse {
	type response struct {
		Success    bool   `json:"success"`
		NewToken   string `json:"new_token"`
		LoginToken string `json:"login_token"`
		Uid        uint   `json:"uid"`
	}

	// check vercode
	vercode_obj, err := auth.ImageVercodeFromToken(vercode_token)
	if err != nil {
		return types.ErrorResponse(-403, "invalid vercode token")
	}
	if !vercode_obj.Check("normal_login", vercode) {
		token, err := vercode_obj.ToToken()
		if err != nil {
			return types.ErrorResponse(-403, "vercode token is expired")
		}
		return types.SuccessResponse(response{
			Success:  false,
			NewToken: token,
		})
	}

	user_obj, err := db.GetOne[types.User](db.Equal("email", email))
	if err != nil {
		return types.ErrorResponse(-403, "invalid email or password")
	}

	password_obj, err := db.GetOne[types.Password](db.Equal("user_id", user_obj.ID))
	if err != nil {
		return types.ErrorResponse(-403, "invalid email or password")
	}

	if !password_obj.Check(math.HashPassword(password)) {
		return types.ErrorResponse(-403, "invalid email or password")
	}

	// generate token
	token, err := auth.RequestLoginToken(password_obj.UserID, time.Hour*24*7)
	if err != nil {
		return types.ErrorResponse(-500, "internal error")
	}

	return types.SuccessResponse(response{
		Success:    true,
		LoginToken: token,
		Uid:        password_obj.UserID,
	})
}

func UserRegister(username, password, vercode, vercode_token, invite_code string) *types.CoshubResponse {
	type response struct {
		Success  bool   `json:"success"`
		NewToken string `json:"new_token"`
	}

	vercode_obj, err := auth.TransportVercodeFromToken(vercode_token)
	if err != nil {
		return types.ErrorResponse(-403, "invalid vercode token")
	}

	if !vercode_obj.Check("normal_register", vercode) {
		if vercode_obj.IsMaxCount() {
			return types.ErrorResponse(-403, "vercode token is expired")
		} else {
			token, err := vercode_obj.ToToken()
			if err != nil {
				return types.ErrorResponse(-403, "vercode token is expired")
			}
			return types.SuccessResponse(response{
				Success:  false,
				NewToken: token,
			})
		}
	}

	email := vercode_obj.Identifier
	if !static.EMAIL_REGX.MatchString(email) {
		return types.ErrorResponse(-403, "invalid email")
	}

	// create user
	_, err = user.CreateUser(username, email, password)
	if err != nil {
		return types.ErrorResponse(-403, "invalid email or password")
	}

	resp := types.SuccessResponse(response{
		Success: true,
	})

	return resp
}

func UserLoginGithub(code string) *types.CoshubResponse {
	type response struct {
		Success    bool   `json:"success"`
		LoginToken string `json:"login_token"`
		Uid        uint   `json:"uid"`
	}

	user, err := auth.LoginWithGithub(code)
	if err != nil {
		return types.ErrorResponse(-403, "invalid github code")
	}

	// generate token
	token, err := auth.RequestLoginToken(user.ID, time.Hour*24*7)
	if err != nil {
		return types.ErrorResponse(-500, "internal error")
	}

	return types.SuccessResponse(response{
		Success:    true,
		LoginToken: token,
		Uid:        user.ID,
	})
}

func UserProfile(user *types.User) *types.CoshubResponse {
	var response struct {
		Username     string            `json:"username"`
		InviteCode   string            `json:"invite_code"`
		TotalSerials uint              `json:"total_serials"`
		TotalBalance uint              `json:"total_balance"`
		Profile      types.UserProfile `json:"profile"`
	}

	response.Username = user.Username
	response.Profile = user.Profile
	return types.SuccessResponse(response)
}

func UserUpdate(user *types.User, username string, avatar string, sign string) *types.CoshubResponse {
	type response struct {
		Success bool `json:"success"`
	}

	user.Username = username
	user.Profile.Avatar = avatar
	user.Profile.Sign = sign

	if model.UpdateUser(user) != nil {
		return types.ErrorResponse(-500, "internal error")
	}

	return types.SuccessResponse(response{
		Success: true,
	})
}

type UserKasumiLoginToken struct {
	Status bool `json:"status"`
	Uid    uint `json:"uid"`
}

func UserRequestLoginKasumiAccept(user *types.User, token string) *types.CoshubResponse {
	type response struct {
		Success bool `json:"success"`
	}

	token_obj := &UserKasumiLoginToken{
		Status: true,
		Uid:    user.ID,
	}

	manager := db.NewCacheManager()
	err := db.CacheSetGeneric(manager, token, token_obj, time.Minute)
	if err != nil {
		return types.ErrorResponse(-500, "internal error")
	}

	return types.SuccessResponse(response{true})
}

func UserRequestLoginKasumiCheck(token string) *types.CoshubResponse {
	type response struct {
		Success bool   `json:"success"`
		Token   string `json:"token"`
		Uid     uint   `json:"uid"`
	}

	manager := db.NewCacheManager()
	token_obj, err := db.CacheGetGeneric[UserKasumiLoginToken](manager, token)
	if err != nil && err != db.ErrCacheKeyNotFount {
		return types.ErrorResponse(-500, "internal error")
	}

	if err == db.ErrCacheKeyNotFount {
		return types.SuccessResponse(response{false, "", 0})
	}

	if token_obj == nil {
		return types.SuccessResponse(response{false, "", 0})
	}

	db.CacheDelGeneric(manager, token, &UserKasumiLoginToken{})

	// generate token
	token, err = auth.RequestLoginToken(token_obj.Uid, time.Hour*24*7)
	if err != nil {
		return types.ErrorResponse(-500, "internal error")
	}

	return types.SuccessResponse(response{true, token, token_obj.Uid})
}
