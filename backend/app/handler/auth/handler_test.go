package auth

import (
	"connectrpc.com/connect"
	"context"
	"go.uber.org/mock/gomock"
	authv1 "sudoku/gen/sudoku/auth/v1"
	"sudoku/service/auth"
	"sudoku/service/auth/mock"
	"testing"
)

func TestHandler_SignIn(t *testing.T) {
	ctrl := gomock.NewController(t)

	type fields struct {
		authService auth.IAuthService
	}
	type args struct {
		ctx context.Context
		req *connect.Request[authv1.SignInRequest]
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantFn  func(t *testing.T, res *connect.Response[authv1.SignInResponse])
		wantErr bool
	}{
		{
			name: "無効なprovider",
			fields: fields{
				authService: mock.NewMockIAuthService(ctrl),
			},
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&authv1.SignInRequest{
					Provider: authv1.OAuthProvider_OAUTH_PROVIDER_UNSPECIFIED,
				}),
			},
			wantFn:  nil,
			wantErr: true,
		},
		{
			name: "正常系",
			fields: fields{
				authService: func() auth.IAuthService {
					s := mock.NewMockIAuthService(ctrl)
					s.EXPECT().SignIn(gomock.Any()).Return(auth.SignInOutput{
						AuthorizationURL: "https://example.com?client_id=123&state=random",
						StateJWT:         "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdGF0ZSI6InJhbmRvbSJ9.Wcs4BFdRKf3tObT4LDaU5frvBU8HOeepNs8JoP3YmY0",
					}, nil)
					return s
				}(),
			},
			args: args{
				ctx: context.Background(),
				req: connect.NewRequest(&authv1.SignInRequest{
					Provider: authv1.OAuthProvider_OAUTH_PROVIDER_GITHUB,
				}),
			},
			wantFn: func(t *testing.T, res *connect.Response[authv1.SignInResponse]) {
				// TODO: ちゃんと検証する
				if res.Msg.AuthorizationUrl == "" {
					t.Error("AuthorizationUrl is empty")
				}
				if res.Header().Get("Set-Cookie") == "" {
					t.Error("Cookie is empty")
				}
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				authService: tt.fields.authService,
			}
			got, err := h.SignIn(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantFn != nil {
				tt.wantFn(t, got)
			}
		})
	}
}
