package auth

import (
	"context"
	"time"

	"github.com/Yeuoly/billboards/internal/db"
	"github.com/Yeuoly/billboards/internal/utils/cocurrent"
	"github.com/Yeuoly/billboards/internal/utils/math"
)

var (
	user_token_cud_lock = cocurrent.GetRandomHighGranularityMutex()
)

/*
	auth controller for all auth related routes
*/

type authMiddlewareType struct {
	Tokens []authMiddlewareToken `json:"tokens"`
	Uid    uint                  `json:"uid"`
}

type authMiddlewareToken struct {
	Token      string    `json:"token"`
	LoginTime  time.Time `json:"login_time"`
	Expiration time.Time `json:"expiration"`
}

func GetUserAuthTokens(uid uint) []authMiddlewareToken {
	cache_manager := db.NewCacheManager()

	auth, err := db.CacheGetGeneric[authMiddlewareType](cache_manager, uid, "auth")
	if err != nil {
		return nil
	}

	// check if token expired, if expired, delete it
	var new_tokens []authMiddlewareToken
	for _, token := range auth.Tokens {
		if token.Expiration.After(time.Now()) {
			new_tokens = append(new_tokens, token)
		}
	}

	// if token expired, update cache
	if len(new_tokens) != len(auth.Tokens) {
		auth.Tokens = new_tokens
		ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
		if err := user_token_cud_lock.Lock(int(uid), ctx); err != nil {
			return auth.Tokens
		}
		defer user_token_cud_lock.Unlock(int(uid))
		db.CacheSetGeneric[authMiddlewareType](cache_manager, uid, auth, time.Hour*24*7, "auth")
	}

	return auth.Tokens
}

// request a login token with expire time
func RequestLoginToken(uid uint, expireTime time.Duration) (string, error) {
	token := math.RandomString(32)
	expiration := time.Now().Add(expireTime)

	cache_manager := db.NewCacheManager()

	auth, err := db.CacheGetGeneric[authMiddlewareType](cache_manager, uid, "auth")
	if err != nil && err != db.ErrCacheKeyNotFount {
		return "", err
	}

	if err == db.ErrCacheKeyNotFount {
		auth = &authMiddlewareType{
			Tokens: []authMiddlewareToken{},
			Uid:    uid,
		}
	}

	auth.Tokens = append(auth.Tokens, authMiddlewareToken{
		Token:      token,
		LoginTime:  time.Now(),
		Expiration: expiration,
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	if err := user_token_cud_lock.Lock(int(uid), ctx); err != nil {
		return "", err
	}
	defer user_token_cud_lock.Unlock(int(uid))

	if err := db.CacheSetGeneric[authMiddlewareType](cache_manager, uid, auth, time.Hour*24*7, "auth"); err != nil {
		return "", err
	}

	return token, nil
}

func CheckLoginToken(uid uint, token string) (bool, uint) {
	tokens := GetUserAuthTokens(uid)
	for _, t := range tokens {
		if t.Token == token {
			return true, uid
		}
	}

	return false, 0
}
