package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"net/url"
	"testing"
)

func TestService_SignIn(t *testing.T) {
	type fields struct {
		config ServiceConfig
	}
	type args struct {
		input SignInInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantFn  func(t *testing.T, got SignInOutput)
		wantErr bool
	}{
		{
			name: "URLとJWTが正しく生成される",
			fields: fields{
				config: ServiceConfig{
					JWTSecret:        "sudoku_secret",
					OAuthClientID:    "client_id",
					OAuthRedirectURI: "http://127.0.0.1:3000/auth/callback",
				},
			},
			args: args{
				input: SignInInput{
					Provider: OAuthProviderGitHub,
				},
			},
			wantFn: func(t *testing.T, got SignInOutput) {
				// AuthorizationURL に必要なパラメータが含まれているか
				parsedURL, err := url.Parse(got.AuthorizationURL)
				if err != nil {
					t.Errorf("failed to parse URL: %v", err)
					return
				}
				params := parsedURL.Query()
				if !(params.Has("client_id") && params.Has("redirect_uri") && params.Has("state")) {
					t.Errorf("missing query parameters: %v", params)
					return
				}
				state := params.Get("state")

				// JWT が正しい State を含むかどうか
				token, err := jwt.ParseWithClaims(got.StateJWT, &StateClaims{}, func(token *jwt.Token) (interface{}, error) {
					return nil, nil
				})
				claims, ok := token.Claims.(*StateClaims)
				if !ok {
					t.Errorf("failed to parse JWT: %v", err)
					return
				}
				if state != claims.State {
					t.Errorf("state mismatch: %v != %v", state, claims.State)
					return
				}
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				config: tt.fields.config,
			}
			got, err := s.SignIn(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.wantFn(t, got)
		})
	}
}
