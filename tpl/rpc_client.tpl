package grpc_client

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	// UserClient user.UserClient
)

func Setup() {
	// UserClient, err = user.NewUserClient(setup("127.0.0.1:8989"))
}


func setup(target string) (*grpc.ClientConn, error) {
	timeout, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return grpc.DialContext(timeout,
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithConnectParams(grpc.ConnectParams{Backoff: backoff.DefaultConfig}),
		grpc.WithChainUnaryInterceptor(
			error2NormalInterceptor(),
			traceInterceptor(),
		))
}


// error2NormalInterceptor 将 GRPC 错误替换为普通错误
func error2NormalInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			if fromError, ok := status.FromError(err); ok {
				err = errors.New(fromError.Message())
			}
		}
		return err
	}
}

// traceInterceptor 添加 traceID
func traceInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.Pairs()
		}

		value := ctx.Value("x-request-id")
			if requestID, ok := value.(string); ok && requestID != "" {
			md["x-request-id"] = []string{requestID}
		}
		return invoker(metadata.NewOutgoingContext(ctx, md), method, req, reply, cc, opts...)
	}
}
