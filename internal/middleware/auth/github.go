package auth

import (
	"errors"
	"fmt"

	"github.com/Yeuoly/billboards/internal/db"
	"github.com/Yeuoly/billboards/internal/db/model"
	"github.com/Yeuoly/billboards/internal/middleware/user"
	"github.com/Yeuoly/billboards/internal/static"
	"github.com/Yeuoly/billboards/internal/types"
	"github.com/Yeuoly/billboards/internal/utils/cocurrent"
	"github.com/Yeuoly/billboards/internal/utils/thirdparty/github"
)

var (
	github_login_lock = cocurrent.GetRandomHighGranularityMutex()
	ErrGithubIdInUse  = errors.New("github id is in use")

	github_client_id     = ""
	github_client_secret = ""
	github_redirect_uri  = ""
	proxy_address        = ""
	proxy_username       = ""
	proxy_password       = ""
)

func LoginWithGithub(code string) (*types.User, error) {
	if github_client_id == "" {
		config := static.GetBocchiGlobalConfigurations()
		github_client_id = config.Github.ClientID
		github_client_secret = config.Github.ClientSecret
		github_redirect_uri = config.Github.RedirectURI
		proxy_address = config.Proxy.Address
		proxy_username = config.Proxy.Username
		proxy_password = config.Proxy.Password
	}

	resp, err := github.GithubRequestAccessToken(
		github_client_id,
		github_client_secret,
		code,
		github_redirect_uri,
		proxy_address,
		proxy_username,
		proxy_password,
	)

	if err != nil {
		return nil, err
	}

	if resp.Err != "" {
		return nil, errors.New(resp.Err)
	}

	if resp.AccessToken == "" {
		return nil, errors.New("access_token is empty")
	}

	gitub_user, err := github.GithubRequestUser(
		resp.AccessToken,
		resp.TokenType,
		proxy_address,
		proxy_username,
		proxy_password,
	)

	if err != nil {
		return nil, err
	}

	github_id := gitub_user.ID
	github_username := gitub_user.LoginName

	if !github_login_lock.TryLock(github_id) {
		return nil, ErrGithubIdInUse
	}
	defer github_login_lock.Unlock(github_id)

	// check if github_id exists
	github_bind, err := db.GetOne[types.UserGithubBind](db.Equal("github_id", github_id))
	if err == nil {
		return model.GetUser(github_bind.UserID)
	}

	// create user
	user, err := user.CreateUser(github_username, fmt.Sprintf("customer_github_%d@miduoduo.org", github_id), "", func(u *types.User) error {
		// create github_bind
		github_bind := &types.UserGithubBind{
			UserID:     u.ID,
			GithubId:   github_id,
			GithubName: github_username,
		}
		return db.Create(github_bind)
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}
