package auth

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

const (
	baseURLGitHub = "https://github.com/login/oauth/access_token"
)

type OAuthClientImpl struct {
	clientID     string
	clientSecret string
}

func NewOAuthClientImpl(clientID string, clientSecret string) *OAuthClientImpl {
	return &OAuthClientImpl{
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

// GetAccessToken GitHub 用の実装
func (c *OAuthClientImpl) GetAccessToken(code string, redirectURI string) (string, error) {
	data := url.Values{}
	data.Set("client_id", c.clientID)
	data.Set("client_secret", c.clientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)

	req, err := http.NewRequest(http.MethodPost, baseURLGitHub, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// Scope, TokenType の使い道は?
	var jsonRes struct {
		AccessToken string `json:"access_token"`
		Scope       string `json:"scope"`
		TokenType   string `json:"token_type"`
	}
	if err := json.NewDecoder(res.Body).Decode(&jsonRes); err != nil {
		return "", err
	}

	return jsonRes.AccessToken, nil
}
