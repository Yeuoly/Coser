package oss

import (
	"io"
	"math/rand"
	"path"
	"strconv"
	"strings"
	"time"

	internal_strings "github.com/Yeuoly/coshub/internal/utils/strings"

	"github.com/Yeuoly/coshub/internal/static"
	"github.com/Yeuoly/coshub/internal/utils/math"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var oss_endpoint string
var oss_bucket_name string
var access_key_id string
var access_key_secret string
var oss_client *oss.Client
var oss_bucket *oss.Bucket

func InitOSS() {
	var err error

	config := static.GetCoshubGlobalConfigurations()

	access_key_id = config.Alioss.AccessKeyId
	access_key_secret = config.Alioss.AccessKeySecret
	oss_endpoint = config.Alioss.Endpoint
	oss_bucket_name = config.Alioss.Bucket

	oss_client, err = oss.New(
		"https://"+oss_endpoint,
		access_key_id,
		access_key_secret,
	)

	if err != nil {
		panic("init oss failed")
	}

	oss_bucket, err = oss_client.Bucket(oss_bucket_name)

	if err != nil {
		panic("init bucket failed")
	}
}

type OssConfig struct {
	typ string
	val interface{}
}

func OssAutoDatePath(crypt string) OssConfig {
	config := OssConfig{
		typ: "auto_date_path",
	}

	if crypt == "" {
		crypt = "md5"
	}

	fn := func(origin_name string) string {
		if origin_name == "" {
			origin_name = strconv.Itoa(int(time.Now().Unix())) + strconv.FormatInt(time.Now().UnixNano(), 16)
		}

		ext := path.Ext(origin_name)

		switch crypt {
		case "md5":
			origin_name = math.Md5(origin_name)
		case "sha1":
			origin_name = math.Sha1(origin_name)
		case "sha256":
			origin_name = math.Sha256(origin_name)
		case "sha512":
			origin_name = math.Sha512(origin_name)
		}
		origin_name += strconv.FormatInt(time.Now().UnixNano(), 36) + strconv.Itoa(int(rand.Int31())) + ext
		origin_name = "upload/" + time.Now().Format("2006/01/02/") + origin_name

		return origin_name
	}

	config.val = fn

	return config
}

func UploadToPublicOss(reader io.Reader, filename string, config ...OssConfig) (string, error) {
	options := []oss.Option{
		oss.ObjectACL(oss.ACLPublicRead),
	}

	for _, v := range config {
		switch v.typ {
		case "auto_date_path":
			fn := v.val.(func(string) string)
			filename = fn(filename)
		}
	}

	err := oss_bucket.PutObject(filename, reader, options...)

	if err != nil {
		return "", err
	}

	return internal_strings.StringJoin(
		"https://",
		oss_bucket_name,
		".",
		oss_endpoint,
		"/",
		filename,
	), nil
}

func UploadToPrivateOss(reader io.Reader, filename string, config ...OssConfig) (string, error) {
	options := []oss.Option{
		oss.ObjectACL(oss.ACLPrivate),
	}

	for _, v := range config {
		switch v.typ {
		case "auto_date_path":
			fn := v.val.(func(string) string)
			filename = fn(filename)
		}
	}

	err := oss_bucket.PutObject(filename, reader, options...)

	if err != nil {
		return "", err
	}

	return internal_strings.StringJoin(
		"https://",
		oss_bucket_name,
		".",
		oss_endpoint,
		"/",
		filename,
	), nil
}

// expires is the duration of the url, such as 30 * time.Minute
func RequestTempraryReadUrl(filename string, expires time.Duration) (string, error) {
	// trim httpsxxxx from filename
	filename = strings.TrimPrefix(filename, internal_strings.StringJoin(
		"https://",
		oss_bucket_name,
		".",
		oss_endpoint,
		"/",
	))

	signed_url, err := oss_bucket.SignURL(filename, oss.HTTPGet, int64(expires.Seconds()))

	if err != nil {
		return "", err
	}

	return signed_url, nil
}

// expires is the duration of the url, such as 30 * time.Minute
func RequestTempraryWriteUrl(filename string, expires time.Duration) (string, error) {
	// trim httpsxxxx from filename
	filename = strings.TrimPrefix(filename, internal_strings.StringJoin(
		"https://",
		oss_bucket_name,
		".",
		oss_endpoint,
		"/",
	))

	signed_url, err := oss_bucket.SignURL(filename, oss.HTTPPut, int64(expires.Seconds()))

	if err != nil {
		return "", err
	}

	return signed_url, nil
}
