package static

import "regexp"

var (
	EMAIL_REGX = regexp.MustCompile(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`)
	PHONE_REGX = regexp.MustCompile(`^1[0-9]{10}$`)
)
