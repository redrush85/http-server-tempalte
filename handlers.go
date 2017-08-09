package main

import (
	"net/http"

	tracer "github.com/opentracing/opentracing-go"
)

// Handler struct
type Handler struct {
	DB      string
	Service Service
}

// Index - index handler
func (h Handler) Index() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		rootSpan, rootCtx := tracer.StartSpanFromContext(r.Context(), "IndexHandler")
		defer rootSpan.Finish()

		w.Write([]byte(h.DB + h.Service.Hello(rootCtx, "alex")))
	}

	return http.HandlerFunc(fn)
}

// Panic - handler with panic
func (h Handler) Panic() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		panic("panic")
	}

	return http.HandlerFunc(fn)
}
