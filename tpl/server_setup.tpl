package server

import (
    "context"
    "net"
    "fmt"

    "{{.pbPkg}}"

    "github.com/google/uuid"
    grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
    "github.com/pkg/errors"
    "github.com/sirupsen/logrus"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/metadata"
    "google.golang.org/grpc/status"
)

func Setup() *grpc.Server {
    rpcServer := grpc.NewServer(
        grpc.ChainUnaryInterceptor(
            grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(logPanic)),
            traceInterceptor(),
        ))
    {{.pbLastPkg}}.Register{{.serverName}}Server(rpcServer, &{{.serverName}}Server{})
    listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "{{.host}}", {{.port}}))
    if err != nil {
        logrus.Panic("%+v", errors.WithStack(err))
    }
    go func() {
        rpcServer.Serve(listen)
    }()
    return rpcServer
}

// logPanic grpc 服务出现 panic 记录日志
func logPanic(p any) error {
    logrus.Errorf("系统错误 %+v", errors.Errorf("%v", p))
    return status.Errorf(codes.Internal, "系统错误: %v", p)
}

// traceInterceptor 添加 traceID
func traceInterceptor() grpc.UnaryServerInterceptor {
    return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
        md, ok := metadata.FromIncomingContext(ctx)
        if !ok {
            md = metadata.Pairs()
        }
        // Set request ID for context.
        requestIDs := md["x-request-id"]
        if len(requestIDs) >= 1 {
            ctx = context.WithValue(ctx, "x-request-id", requestIDs[0])
            return handler(ctx, req)
        }

        // Generate request ID and set context if not exists.
        requestID := uuid.New().String()
        ctx = context.WithValue(ctx, "x-request-id", requestID)

        return handler(ctx, req)
    }
}
