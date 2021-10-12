package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"net/http"

	"github.com/pardot/oidc/core"
)

const (
	sessIDCookie = "sessID"

	upstreamAllowQuery = "data.upstream.allow"
)

type server struct {
	issuer          string
	oidcsvr         *core.OIDC
	providers       map[string]Provider
	storage         Storage
	tokenValidFor   time.Duration
	refreshValidFor time.Duration

	// upstreamPolicy is rego code applied to claims from upstream IDP
	upstreamPolicy []byte
}

func (s *server) authorization(w http.ResponseWriter, req *http.Request) {
	ar, err := s.oidcsvr.StartAuthorization(w, req)
	if err != nil {
		log.Printf("error starting authorization: %v", err)
		return
	}

	w.Write([]byte(`<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>title</title>
    <link rel="stylesheet" href="style.css">
    <script src="script.js"></script>
  </head>
  <body>`))

	for id, p := range s.providers {
		panel, err := p.LoginPanel(req, ar)
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, `<div id="%s-panel">`, id)
		w.Write([]byte(panel))
		w.Write([]byte(`</div>`))
	}

	w.Write([]byte(`</body>
	</html>`))
}

func (s *server) token(w http.ResponseWriter, req *http.Request) {
	err := s.oidcsvr.Token(w, req, func(tr *core.TokenRequest) (*core.TokenResponse, error) {
		auth, ok, err := s.storage.GetAuthentication(req.Context(), tr.SessionID)
		if err != nil {
			return nil, fmt.Errorf("getting authentication for session %s", tr.SessionID)
		}
		if !ok {
			return nil, fmt.Errorf("no authentication for session %s", tr.SessionID)
		}

		idt := tr.PrefillIDToken(s.issuer, auth.Subject, time.Now().Add(s.tokenValidFor))

		// oauth2 proxy wants this, when we don't have useinfo
		// TODO - scopes/userinfo etc.
		idt.Extra["email"] = auth.EMail
		idt.Extra["email_verified"] = true

		return &core.TokenResponse{
			AccessTokenValidUntil:  time.Now().Add(s.tokenValidFor),
			RefreshTokenValidUntil: time.Now().Add(s.refreshValidFor),
			IssueRefreshToken:      tr.SessionRefreshable, // always allow it if we want it
			IDToken:                idt,
		}, nil
	})
	if err != nil {
		log.Printf("error in token endpoint: %v", err)
	}
}

func (s *server) AddHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/auth", s.authorization)
	mux.HandleFunc("/token", s.token)
}

type authSessionManager struct {
	storage Storage
	oidcsvr *core.OIDC
}

func (a *authSessionManager) GetMetadata(ctx context.Context, sessionID string, into interface{}) (ok bool, err error) {
	md, ok, err := a.storage.GetMetadata(ctx, sessionID)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}
	if err := json.Unmarshal(md.ProviderMetadata, into); err != nil {
		return false, fmt.Errorf("unmarshaling provider metadata: %v", err)
	}
	return true, nil
}

func (a *authSessionManager) PutMetadata(ctx context.Context, sessionID string, d interface{}) error {
	md, ok, err := a.storage.GetMetadata(ctx, sessionID)
	if err != nil {
		return err
	}
	if !ok {
		md = Metadata{}
	}
	md.ProviderMetadata, err = json.Marshal(d)
	if err != nil {
		return fmt.Errorf("marshaling provider metadata: %v", err)
	}
	return nil
}

func (a *authSessionManager) Authenticate(w http.ResponseWriter, req *http.Request, sessionID string, auth Authentication) {
	if err := a.storage.Authenticate(req.Context(), sessionID, auth); err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}
	// TODO - we need to fill this. This is likely going to need information
	// about the provider (acr), requested claims, etc. This probably goes in
	// the server metadata field
	az := &core.Authorization{
		Scopes: []string{"openid"},
	}
	a.oidcsvr.FinishAuthorization(w, req, sessionID, az)
}
