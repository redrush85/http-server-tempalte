package main

import (
	"context"
	"time"

	"go.uber.org/zap"

	tracer "github.com/opentracing/opentracing-go"
)

// Service interface
type Service interface {
	Hello(ctx context.Context, name string) string
}

// AppService - Service interface implementation
type AppService struct{}

// Hello - returns Hello + name string
func (a *AppService) Hello(ctx context.Context, name string) string {
	childSpan, _ := tracer.StartSpanFromContext(ctx, "service_hello")
	defer childSpan.Finish()

	logger.Info("This is the coolest log",
		zap.String("url", "http://127.0.0.1/"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	return "Hello " + name
}
