package svc

import (
	"context"
	"time"
)

type GrpcContext struct {
	grpcCtx context.Context
}

func NewGrpcContext(grpcCtx context.Context) *GrpcContext {
	return &GrpcContext{
		grpcCtx: grpcCtx,
	}
}

func (c *GrpcContext) Deadline() (deadline time.Time, ok bool) {
	return c.grpcCtx.Deadline()
}

func (c *GrpcContext) Done() <-chan struct{} {
	return c.grpcCtx.Done()
}

func (c *GrpcContext) Err() error {
	return c.grpcCtx.Err()
}

func (c *GrpcContext) Value(key any) any {
	return c.grpcCtx.Value(key)
}
