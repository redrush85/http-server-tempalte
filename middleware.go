package main

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	tracing "github.com/opentracing/opentracing-go"
)

// RecoverMiddleware - recover http handler from panics
func RecoverMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				logger.Error("Request recovered from panic.", zap.Any("error", r))
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	})
}

// OpenTracing middleware adds support for OpenTracing protocol:
// 1. It extracts previous span from request headers;
// 2. It wraps every request into it's own tracing span.
func OpenTracing(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {

		span, spanCtx := createRequestSpan(r)
		defer span.Finish()

		next.ServeHTTP(w, r.WithContext(spanCtx))
	}

	return http.HandlerFunc(fn)
}

// createRequestSpan is a helper function that creates request tracing span.
func createRequestSpan(r *http.Request) (span tracing.Span, spanCtx context.Context) {

	var sp tracing.Span

	tracingSpanCtx, err := tracing.GlobalTracer().Extract(tracing.TextMap, r.Header)
	if err != nil {
		sp = tracing.StartSpan(ServiceName)
	} else {
		sp = tracing.StartSpan(ServiceName, tracing.ChildOf(tracingSpanCtx))
	}

	return sp, tracing.ContextWithSpan(r.Context(), sp)
}
