package auth

import (
	"encoding/json"
	"time"

	"github.com/Yeuoly/billboards/internal/utils/math"
	"github.com/Yeuoly/billboards/internal/utils/typ"
)

type TransportVercode struct {
	TextVercode
	Identifier    string `json:"identifier"`
	TransportType int    `json:"transport_type"`
}

const (
	TRANSPORT_TYPE_EMAIL = iota + 1
	TRANSPORT_TYPE_PHONE
)

func NewEmailVercode(identifier string, result string, maxCount int, method string) *TransportVercode {
	return &TransportVercode{
		Identifier:    identifier,
		TransportType: TRANSPORT_TYPE_EMAIL,
		TextVercode: TextVercode{
			Magic:    math.RandomString(8),
			Result:   result,
			StartAt:  time.Now(),
			ExpireAt: time.Now().Add(time.Minute * 5),
			Count:    0,
			MaxCount: maxCount,
			Method:   method,
		},
	}
}

func (i *TransportVercode) ToToken() (string, error) {
	text, err := json.Marshal(i)
	if err != nil {
		return "", err
	}

	return math.AesEncrypt(typ.BytesToString(text), global_vercode_key, global_vercode_iv)
}

func TransportVercodeFromToken(token string) (*TransportVercode, error) {
	text, err := math.AesDecrypt(token, global_vercode_key, global_vercode_iv)
	if err != nil {
		return nil, err
	}

	var result TransportVercode
	err = json.Unmarshal(typ.StringTobytes(text), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
