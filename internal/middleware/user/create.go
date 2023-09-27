package user

import (
	"errors"

	"github.com/Yeuoly/billboards/internal/db"
	"github.com/Yeuoly/billboards/internal/types"
	"github.com/Yeuoly/billboards/internal/utils/math"
	"gorm.io/gorm"
)

/*
Create User, username and email are unique and required
password is optional, if password is empty, then the user will be created without password
success is optional, if success is not empty, then it will be called after the user is created
but if success return an error, the transaction will be rollback
*/
func CreateUser(username, email, password string, success ...func(*types.User) error) (*types.User, error) {
	var p_user *types.User
	return p_user, db.WithTransaction(func(tx *gorm.DB) error {
		// check username and email
		if username == "" {
			return errors.New("username is required")
		} else if email == "" {
			return errors.New("email is required")
		}

		if _, err := db.GetOne[types.User](db.Equal("email", email)); err == nil {
			return errors.New("email is already exists")
		}

		user := &types.User{
			Username: username,
			Email:    email,
			Role:     types.USER_ROLE_USER,
		}
		if err := db.Create(user, tx); err != nil {
			return err
		}

		p_user = user

		if password != "" {
			password_obj := &types.Password{
				UserID:   user.ID,
				Password: math.HashPassword(password),
			}
			if err := db.Create(password_obj, tx); err != nil {
				return err
			}
		}

		profile := &types.UserProfile{
			UserID: user.ID,
			Avatar: "",
			Sign:   "这个人很懒，什么都没有留下",
		}

		if err := db.Create(profile, tx); err != nil {
			return err
		}

		if len(success) > 0 {
			return success[0](user)
		}

		return nil
	})
}
