package session

import (
	"context"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

type contextKey string

var contextKeySession = contextKey("session")

// ContextSession is chi middleware that will register the given session on the context.
func ContextSession(s sessions.Store, sessionName string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			session, err := s.Get(r, sessionName)
			if err != nil { // ignore bad key, just start new session
				mErr, ok := err.(securecookie.MultiError)
				if !ok || mErr.Error() != securecookie.ErrMacInvalid.Error() {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}

			r = r.WithContext(context.WithValue(r.Context(), contextKeySession, session))

			// hook in to response, so we can save the session last minute
			hw := &hookRW{
				ResponseWriter: w,
				req:            r,
				sess:           session,
			}

			next.ServeHTTP(hw, r)
		}
		return http.HandlerFunc(fn)
	}
}

// FromContext returns the session instance from the given context
func FromContext(ctx context.Context) *sessions.Session {
	s, ok := ctx.Value(contextKeySession).(*sessions.Session)
	if !ok {
		return nil
	}
	return s
}

type hookRW struct {
	http.ResponseWriter
	req       *http.Request
	sess      *sessions.Session
	sessSaved bool
}

func (h *hookRW) Write(b []byte) (int, error) {
	if !h.sessSaved {
		h.sess.Save(h.req, h.ResponseWriter)
	}
	return h.ResponseWriter.Write(b)
}