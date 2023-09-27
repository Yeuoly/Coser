package auth

import (
	"errors"
	"strconv"

	"github.com/Yeuoly/billboards/internal/static"
	"github.com/Yeuoly/billboards/internal/utils/math"
	email "github.com/Yeuoly/billboards/internal/utils/thirdparty/feishu"
)

var (
	emailManager *email.BocchiEmailManager
)

// SendVercodeEmail sends verification code to user's email and return the encrypted vercode
func SendVercodeEmail(to string, text string, method string) (string, error) {
	if emailManager == nil {
		config := static.GetBocchiGlobalConfigurations()
		emailManager = email.NewBocchiEmailManager(
			config.Email.Host,
			config.Email.Port,
			config.Email.User,
			config.Email.Pass,
		)
	}

	result := strconv.Itoa(math.Random(100000, 999999))
	vercode := NewEmailVercode(to, result, 3, method)

	err := emailManager.SendMail(to, "Bocchi Verification Code", text+"\n Your verification code is: "+result+"\n\n This code will expire in 5 minutes.")
	if err != nil {
		return "", err
	}

	return vercode.ToToken()
}

// CheckEmailVercode checks if the email verification code is correct, and returns the email address
func CheckEmailVercode(token string, result string, method string) (string, error) {
	vercode, err := TransportVercodeFromToken(token)
	if err != nil {
		return "", err
	}

	if !vercode.Check(method, result) {
		return "", errors.New("invalid vercode")
	}

	return vercode.Identifier, nil
}
