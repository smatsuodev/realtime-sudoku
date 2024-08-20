package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sudoku/service/auth"
)

const (
	apiBase    = "https://api.github.com"
	apiVersion = "2022-11-28"
)

var _ auth.GitHubAPI = (*GitHubAPIImpl)(nil)

type GitHubAPIImpl struct{}

func NewGitHubAPIImpl() *GitHubAPIImpl {
	return &GitHubAPIImpl{}
}

func (g *GitHubAPIImpl) GetUser(accessToken string) (*auth.GitHubUser, error) {
	req, err := http.NewRequest(http.MethodGet, apiBase+"/user", nil)
	if err != nil {
		return &auth.GitHubUser{}, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", apiVersion)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &auth.GitHubUser{}, err
	}
	defer res.Body.Close()

	// id 以外も取りたくなるかも
	// refs: https://docs.github.com/ja/rest/users/users?apiVersion=2022-11-28#get-the-authenticated-user
	var jsonRes struct {
		ID    uint   `json:"id"`
		Login string `json:"login"`
	}
	if err := json.NewDecoder(res.Body).Decode(&jsonRes); err != nil {
		return &auth.GitHubUser{}, err
	}

	return &auth.GitHubUser{
		ID:    jsonRes.ID,
		Login: jsonRes.Login,
	}, nil
}
