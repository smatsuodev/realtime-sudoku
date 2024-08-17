package auth

import (
	"net/url"
	"strings"
)

const (
	baseURLGitHub = "https://github.com/login/oauth/authorize"
)

type generateURLParams struct {
	redirectURI string
	state       string
	scope       []string
	clientID    string
}

func generateGitHubAuthURL(params generateURLParams) string {
	p := url.Values{}
	p.Add("client_id", params.clientID)
	p.Add("redirect_uri", params.redirectURI)
	p.Add("state", params.state)
	p.Add("scope", strings.Join(params.scope, " "))

	return baseURLGitHub + "?" + p.Encode()
}
