package main

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"go.uber.org/zap"
)

type loggerCtxKey struct{}
type requestIDCtxKey struct{}
type sessionStoreCtxKey struct{}

// baseMiddleware should wrap all requests to the service
func baseMiddleware(wrapped http.Handler,
	logger *zap.SugaredLogger,
	scHashKey []byte,
	scEncryptKey []byte,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		st := time.Now()

		rid := uuid.New()
		ctx = context.WithValue(ctx, requestIDCtxKey{}, rid)

		sl := logger.With("request_id", rid)
		ctx = context.WithValue(ctx, loggerCtxKey{}, rid)

		store := sessions.NewCookieStore(scHashKey, scEncryptKey)
		ctx = context.WithValue(ctx, sessionStoreCtxKey{}, store)

		ww := &wrapResponseWriter{ResponseWriter: w}

		wrapped.ServeHTTP(ww, r.WithContext(ctx))

		sl.With(
			"method", r.Method,
			"path", r.URL.Path,
			"status", ww.Status(),
			"duration", time.Since(st),
		).Info()
	})
}

func loggerFromContext(ctx context.Context) *zap.SugaredLogger {
	l, ok := ctx.Value(loggerCtxKey{}).(*zap.SugaredLogger)
	if ok {
		return l
	}
	return zap.NewNop().Sugar()
}

func sessionStoreFromContext(ctx context.Context) sessions.Store {
	return ctx.Value(sessionStoreCtxKey{}).(sessions.Store)
}

// httpErrHandler renders out nicer errors
type httpErrHandler struct {
}

func (h *httpErrHandler) Error(w http.ResponseWriter, r *http.Request, err error) {
	l := loggerFromContext(r.Context())
	l.Error(err)
	http.Error(w, "Internal Error", http.StatusInternalServerError)
}

func (h *httpErrHandler) BadRequest(w http.ResponseWriter, r *http.Request, message string) {
	http.Error(w, message, http.StatusBadRequest)
}

func (h *httpErrHandler) Forbidden(w http.ResponseWriter, r *http.Request, message string) {
	http.Error(w, message, http.StatusForbidden)
}

type wrapResponseWriter struct {
	http.ResponseWriter
	st int
	wh bool
}

func (w *wrapResponseWriter) Status() int {
	return w.st
}

func (w *wrapResponseWriter) WriteHeader(code int) {
	if w.wh {
		return
	}
	w.st = code
	w.ResponseWriter.WriteHeader(code)
	w.wh = true
}