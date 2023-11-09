package model

import (
	"time"

	"github.com/Yeuoly/coshub/internal/db"
	"github.com/Yeuoly/coshub/internal/types"
	"gorm.io/gorm"
)

/*
 * Get UserByUid global function, only accept uid as parameter
 */
func GetUser(uid uint) (*types.User, error) {
	cache_manager := db.NewCacheManager()
	p_user, err := db.CacheGetGeneric[types.User](cache_manager, uid, "user")
	if err != nil && err != db.ErrCacheKeyNotFount {
		return nil, err
	}

	if err == nil && p_user != nil {
		return p_user, nil
	}

	var user types.User
	err = db.Run(db.Action(func(tx *gorm.DB) {
		tx.Model(&types.User{}).Preload("UserProfile").First(&user, uid)
	}))

	if err != nil {
		return nil, err
	}

	db.CacheSetGeneric[types.User](cache_manager, user.ID, &user, time.Hour*24, "user")

	return &user, nil
}

/*
 * Update User global function, only accept user as parameter
 */
func UpdateUser(user *types.User) error {
	return db.WithTransaction(func(tx *gorm.DB) error {
		err := db.Update(user)
		if err != nil {
			return err
		}

		err = db.Update(&user.Profile)
		if err != nil {
			return err
		}

		cache_manager := db.NewCacheManager()
		db.CacheDelGeneric(cache_manager, user.ID, user, "user")

		return nil
	})
}

/*
 * Delete User global function, only accept user as parameter
 */
func DeleteUser(user *types.User) error {
	return db.WithTransaction(func(tx *gorm.DB) error {
		err := db.Delete(&user.Profile)
		if err != nil {
			return err
		}

		cache_manager := db.NewCacheManager()
		db.CacheDelGeneric(cache_manager, user.ID, user, "user")

		return nil
	})
}
