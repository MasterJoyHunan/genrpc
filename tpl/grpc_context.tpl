package svc

import (
    "context"
    "time"
)

type GrpcContext struct {
    grpcCtx context.Context
}

func NewGrpcContext(grpcCtx context.Context) *ServiceContext {
    return &GrpcContext{
        grpcCtx: grpcCtx,
    }
}

func (c *ServiceContext) Deadline() (deadline time.Time, ok bool) {
    return c.grpcCtx.Deadline()
}

func (c *ServiceContext) Done() <-chan struct{} {
    return c.grpcCtx.Done()
}

func (c *ServiceContext) Err() error {
    return c.grpcCtx.Err()
}

func (c *ServiceContext) Value(key any) any {
    return c.grpcCtx.Value(key)
}
