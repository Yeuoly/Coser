package auth

import (
	"encoding/json"
	"time"

	"github.com/Yeuoly/billboards/internal/db"
	"github.com/Yeuoly/billboards/internal/utils/math"
	"github.com/Yeuoly/billboards/internal/utils/typ"
)

var (
	global_vercode_key = math.RandomBytes(32)
	global_vercode_iv  = math.RandomBytes(16)
)

// TextVercode is a struct for image vercode, expire time is 5 minutes.
type TextVercode struct {
	Magic    string    `json:"magic"`
	Result   string    `json:"result"`
	StartAt  time.Time `json:"start_at"`
	ExpireAt time.Time `json:"expire_at"`
	Count    int       `json:"count"`
	MaxCount int       `json:"max_count"`
	Method   string    `json:"method"` // login, register, reset_password, change_phone_number etc.
}

func NewImageVercode(result string, maxCount int, method string) *TextVercode {
	return &TextVercode{
		Magic:    math.RandomString(8),
		Result:   result,
		StartAt:  time.Now(),
		ExpireAt: time.Now().Add(time.Minute * 5),
		Count:    0,
		MaxCount: maxCount,
		Method:   method,
	}
}

func (i *TextVercode) IsExpired() bool {
	return time.Now().After(i.ExpireAt)
}

func (i *TextVercode) IsMaxCount() bool {
	return i.Count >= i.MaxCount
}

func (i *TextVercode) Check(method string, result string) bool {
	cache_manager := db.NewCacheManager()
	// check if token is already used
	err := cache_manager.SetNX(cache_manager.GetRedisKey("vercode", i.Magic), 1, time.Minute*5)
	if err != nil {
		return false
	}

	i.Magic = math.RandomString(8)
	success := !i.IsExpired() && !i.IsMaxCount() && i.Result == result && i.Method == method
	i.AddCount()
	return success
}

func (i *TextVercode) AddCount() {
	i.Count++
}

func (i *TextVercode) ToToken() (string, error) {
	text, err := json.Marshal(i)
	if err != nil {
		return "", err
	}

	return math.AesEncrypt(typ.BytesToString(text), global_vercode_key, global_vercode_iv)
}

func ImageVercodeFromToken(token string) (*TextVercode, error) {
	text, err := math.AesDecrypt(token, global_vercode_key, global_vercode_iv)
	if err != nil {
		return nil, err
	}

	var result TextVercode
	err = json.Unmarshal(typ.StringTobytes(text), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
