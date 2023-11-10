package http_service

import (
	"strings"
	"time"

	"github.com/Yeuoly/coshub/internal/db"
	"github.com/Yeuoly/coshub/internal/types"
	"github.com/Yeuoly/coshub/internal/utils/oss"
)

func UploadFile(filename string, content_type string, ip string) *types.CoshubResponse {
	type response struct {
		Url string `json:"url"`
	}

	url, err := oss.RequestTempraryWriteUrl(filename, time.Minute*5, content_type, oss.OssAutoDatePath("md5"))
	if err != nil {
		return types.ErrorResponse(-500, "internal error")
	}

	file := &types.File{
		Url: url,
		Ip:  ip,
	}

	err = db.Create(file)
	if err != nil {
		return types.ErrorResponse(-500, "internal error")
	}

	return types.SuccessResponse(response{
		Url: url,
	})
}

func FinishUploadFile(url string) *types.CoshubResponse {
	type response struct {
		Success bool `json:"success"`
	}

	// get filename
	idx := strings.Index(url, "?")
	if idx != -1 {
		url = url[:idx]
	}
	// delete protocol and domain
	idx = strings.Index(url, "//")
	if idx != -1 {
		url = url[idx+2:]
	}
	idx = strings.Index(url, "/")
	if idx != -1 {
		url = url[idx+1:]
	}

	filename := url

	err := oss.SetObjectACLPublicRead(filename)

	if err != nil {
		return types.ErrorResponse(-500, "internal error")
	}

	return types.SuccessResponse(response{
		Success: true,
	})
}
