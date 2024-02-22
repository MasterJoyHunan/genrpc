package svc

import (
	"context"
	"time"
	// "github.com/sirupsen/logrus"
)

type GrpcContext struct {
	// Log     *logrus.Entry
	grpcCtx context.Context
}

func NewGrpcContext(grpcCtx context.Context) *ServiceContext {
	// value := grpcCtx.Value(util.TrafficKey)
	// traceId := ""
	// if value != nil {
	//    traceId = value.(string)
	// }
	return &GrpcContext{
		grpcCtx: grpcCtx,
		// Log:     logrus.WithField(util.TraceID, traceId),
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
