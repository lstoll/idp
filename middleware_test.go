package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
)

func TestMiddlewareSessions(t *testing.T) {
	smgr := scs.New()
	smgr.Lifetime = 1 * time.Minute

	mux := http.NewServeMux()

	mux.HandleFunc("/set", func(_ http.ResponseWriter, req *http.Request) {
		key := req.URL.Query().Get("key")
		if key == "" {
			t.Fatal("query with no key")
		}

		value := req.URL.Query().Get("value")
		if key == "" {
			t.Fatal("query with no value")
		}

		smgr.Put(req.Context(), key, value)
	})

	mux.HandleFunc("/get", func(rw http.ResponseWriter, req *http.Request) {
		key := req.URL.Query().Get("key")
		if key == "" {
			t.Fatal("query with no key")
		}

		val := smgr.GetString(req.Context(), key)

		if val == "" {
			http.Error(rw, fmt.Sprintf("key %s has no value", key), http.StatusNotFound)
			return
		}

		rw.Write([]byte(val))
	})

	svr := baseMiddleware(mux, smgr)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/set?key=test1&value=value1", nil)

	svr.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("set returned non-200: %d", rec.Code)
	}

	cookies := rec.Result().Cookies()

	rec = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/get?key=test1", nil)
	for _, c := range cookies {
		req.AddCookie(c)
	}

	svr.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("set returned non-200: %d", rec.Code)
	}

	if body, err := io.ReadAll(rec.Body); err != nil && string(body) != "value1" {
		t.Fatalf("wanted response body value1, got %s (err: %v)", string(body), err)
	}
}