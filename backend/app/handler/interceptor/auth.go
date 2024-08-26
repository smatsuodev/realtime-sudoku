package interceptor

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"log"
	commonv1 "sudoku/gen/sudoku/common/v1"
	authH "sudoku/handler/auth"
	authS "sudoku/service/auth"
)

// TODO: streaming の場合は別の interceptor が必要かも
func NewAuthInterceptor(authService authS.IAuthService) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			methodDesc, ok := req.Spec().Schema.(protoreflect.MethodDescriptor)
			if !ok {
				return nil, connect.NewError(connect.CodeInternal, errors.New("invalid method descriptor"))
			}
			opts := methodDesc.Options().(*descriptorpb.MethodOptions)
			requireAuth, ok := proto.GetExtension(opts, commonv1.E_RequireAuth).(bool)
			if !ok {
				return nil, connect.NewError(connect.CodeInternal, errors.New("invalid method option"))
			}

			log.Println("requireAuth", requireAuth)

			if requireAuth {
				sessionCookie := req.Header().Get(authH.SessionCookieName)
				if sessionCookie == "" {
					return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("missing session cookie"))
				}
				sessionID, err := uuid.Parse(sessionCookie)
				if err != nil {
					return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("invalid session cookie"))
				}

				output, err := authService.ValidateSession(authS.ValidateSessionInput{
					SessionID: sessionID,
				})
				if err != nil {
					return nil, connect.NewError(connect.CodeInternal, err)
				}

				if !output.IsValid {
					return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("invalid session"))
				}
			}

			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
