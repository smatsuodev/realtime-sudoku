package config

import (
	"github.com/caarlos0/env/v11"
	"log"
)

type EnvConfig struct {
	JWTSecret          string `env:"JWT_SECRET,required"`
	GitHubClientID     string `env:"GITHUB_CLIENT_ID,required"`
	GitHubClientSecret string `env:"GITHUB_CLIENT_SECRET,required"`
	OAuthRedirectURI   string `env:"OAUTH_REDIRECT_URI,required"`
}

func NewEnvConfig() EnvConfig {
	config, err := env.ParseAs[EnvConfig]()
	if err != nil {
		log.Fatalln(err)
	}
	return config
}
