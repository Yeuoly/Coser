package github

import (
	"fmt"

	"github.com/Yeuoly/billboards/internal/utils/net"
)

type GithubRequestAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	Err         string `json:"error"`
}

type GithubRequestUserResponse struct {
	LoginName string `json:"login"`
	ID        int    `json:"id"`
	NodeID    string `json:"node_id"`
	AvatarURL string `json:"avatar_url"`
	Err       string `json:"error"`
}

type GithubRequestEmailResponse struct {
	Email    string `json:"email"`
	Verified bool   `json:"verified"`
	Primary  bool   `json:"primary"`
}

type GithubRequestUserEmailsResponse []GithubRequestEmailResponse

func GithubRequestAccessToken(
	client_id string,
	client_secret string,
	code string,
	redirect_uri string,
	proxy_address string, proxy_username string, proxy_password string,
) (GithubRequestAccessTokenResponse, error) {
	resp, err := net.SendPostAndParse[GithubRequestAccessTokenResponse](
		"https://github.com/login/oauth/access_token",
		net.HttpHeader(map[string]string{
			"Accept": "application/json",
		}),
		net.HttpProxy(proxy_address, proxy_username, proxy_password),
		net.HttpWithRandomUA(),
		net.HttpPayload(map[string]string{
			"client_id":     client_id,
			"client_secret": client_secret,
			"code":          code,
			"redirect_uri":  redirect_uri,
		}),
	)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

func GithubRequestUser(
	access_token string,
	token_type string,
	proxy_address string, proxy_username string, proxy_password string,
) (GithubRequestUserResponse, error) {
	resp, err := net.SendGetAndParse[GithubRequestUserResponse](
		"https://api.github.com/user",
		net.HttpWithRandomUA(),
		net.HttpProxy(proxy_address, proxy_username, proxy_password),
		net.HttpHeader(map[string]string{
			"Accept":        "application/json",
			"Authorization": fmt.Sprintf("%s %s", token_type, access_token),
		}),
	)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
