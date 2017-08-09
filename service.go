package main

import (
	"context"

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
	childSpan, _ := tracer.StartSpanFromContext(ctx, "my-child-operation")
	defer childSpan.Finish()

	return "Hello " + name
}
