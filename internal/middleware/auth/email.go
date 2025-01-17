package auth

import (
	"errors"
	"strconv"

	"github.com/Yeuoly/coshub/internal/static"
	"github.com/Yeuoly/coshub/internal/utils/math"
	email "github.com/Yeuoly/coshub/internal/utils/thirdparty/feishu"
)

var (
	emailManager *email.BillboardsEmailManager
)

// SendVercodeEmail sends verification code to user's email and return the encrypted vercode
func SendVercodeEmail(to string, text string, method string) (string, error) {
	if emailManager == nil {
		config := static.GetCoshubGlobalConfigurations()
		emailManager = email.NewBillboardsEmailManager(
			config.Email.Host,
			config.Email.Port,
			config.Email.User,
			config.Email.Pass,
		)
	}

	result := strconv.Itoa(math.Random(100000, 999999))
	vercode := NewEmailVercode(to, result, 3, method)

	err := emailManager.SendMail(to, "Billboards Verification Code", text+"\n Your verification code is: "+result+"\n\n This code will expire in 5 minutes.")
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
