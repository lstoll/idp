package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"
	"github.com/justinas/nosurf"
	"github.com/lstoll/cookiesession"
	"github.com/lstoll/oidc/core"
	"github.com/lstoll/oidc/discovery"
	oidcm "github.com/lstoll/oidc/middleware"
	"github.com/oklog/run"
	"golang.org/x/sys/unix"
)

//go:embed web/public/*
var staticFiles embed.FS

func main() {
	debug := flag.Bool("debug", false, "Enable debug logging")

	addr := flag.String("http", "127.0.0.1:8085", "Run the IDP server on the given host:port.")
	configFile := flag.String("config", "config.json", "Path to the config file.")
	enroll := flag.Bool("enroll", false, "Enroll a user into the system.")
	email := flag.String("email", "", "Email address for the user.")
	fullname := flag.String("fullname", "", "Full name of the user.")
	activate := flag.Bool("activate", false, "Activate an enrolled user.")
	userID := flag.String("user-id", "", "ID of user to activate.")

	flag.Parse()

	b, err := os.ReadFile(*configFile)
	if err != nil {
		fatalf("read config file: %v", err)
	}
	var cfg config
	if err := loadConfig(b, &cfg); err != nil {
		fatalf("load config file: %v", err)
	}

	var level slog.Leveler
	if *debug {
		level = slog.LevelDebug
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: level})))

	ctx := context.Background()

	db, err := openDB(cfg.Database)
	if err != nil {
		fatalf("open database at %s: %v", cfg.Database, err)
	}

	if sqlfile := cfg.SQLDatabase; sqlfile != "" {
		sqldb, err := newStorage(ctx, fmt.Sprintf("file:%s?cache=shared&mode=rwc&_journal_mode=WAL", sqlfile))
		if err != nil {
			fatalf("open sqlite database: %v", err)
		}
		if err := migrateSQLToJSON(sqldb, db); err != nil {
			fatalf("migrate SQLite database %s to %s: %v", sqlfile, cfg.Database, err)
		}
	}

	if *enroll {
		if *email == "" {
			fatal("required flag missing: email")
		}
		if *fullname == "" {
			fatal("required flag missing: fullname")
		}

		user, err := db.CreateUser(User{
			Email:     *email,
			FullName:  *fullname,
			Activated: false,
		})
		if err != nil {
			fatalf("create user: %v", err)
		}
		reloadDB(*addr)
		fmt.Printf("Enroll at: %s\n", registrationURL(cfg.Issuer[0].URL, user))
		return
	} else if *activate {
		if *userID == "" {
			fatal("required flag missing: user-id")
		}
		if err := activateUser(db, *userID); err != nil {
			fatalf("ativate user: %v", err)
		}
		reloadDB(*addr)
		return
	}

	if *addr == "" {
		fatal("required flag missing: http")
	}

	issuer := cfg.Issuer[0]

	ks, err := newDerivedKeyset(cfg.EncryptionKey, cfg.PrevEncryptionKey...)
	if err != nil {
		fatalf("derive keyset: %v", err)
	}
	if err := serve(ctx, db, ks, issuer, *addr); err != nil {
		fatalf("start server: %v", err)
	}
}

func serve(ctx context.Context, db *DB, appKeys *derivedKeyset, issuer issuerConfig, addr string) error {
	ksm, err := NewKeysetManager(db)
	if err != nil {
		return fmt.Errorf("creating OIDC keyset manager: %w", err)
	}

	oidcmd := discovery.ProviderMetadata{
		Issuer:                issuer.URL.String(),
		JWKSURI:               issuer.URL.String() + "/keys",
		AuthorizationEndpoint: issuer.URL.String() + "/auth",
		TokenEndpoint:         issuer.URL.String() + "/token",
	}
	oidcHandles := ksm.Handles(KeysetOIDC)
	keysh, err := discovery.NewKeysHandler(oidcHandles.PublicHandle, 1*time.Hour)
	if err != nil {
		return fmt.Errorf("creating discovery keys handler: %w", err)
	}
	discoh, err := discovery.NewConfigurationHandler(&oidcmd, discovery.WithCoreDefaults())
	if err != nil {
		return fmt.Errorf("configuring metadata handler: %w", err)
	}

	clients := &multiClients{
		sources: []core.ClientSource{
			&staticClients{
				clients: issuer.Client,
			},
		},
	}

	oidcsvr, err := core.New(&core.Config{
		AuthValidityTime: 5 * time.Minute,
		CodeValidityTime: 5 * time.Minute,
	}, db.SessionManager(), clients, oidcHandles.Handle)
	if err != nil {
		return fmt.Errorf("failed to create OIDC server instance: %w", err)
	}

	wsSessKeys := &cookiesession.StaticKeys{Encryption: appKeys.webSessCurr, Decryption: appKeys.webSessPrev}
	webSessMgr, err := cookiesession.New[webSession](wsSessKeys, cookiesession.Options{
		MaxAge: 0, // Scopes it to browser lifecycle, which I think is good for now
		Path:   "/",
	})
	if err != nil {
		return fmt.Errorf("creating cookie session for webauthn: %w", err)
	}

	oidcmSessKeys := &cookiesession.StaticKeys{Encryption: appKeys.oidcmSessCurr, Decryption: appKeys.oidcmSessPrev}
	oidcmSessMgr, err := cookiesession.New[oidcMiddlewareSession](oidcmSessKeys, cookiesession.Options{
		MaxAge: 0, // Scopes it to browser lifecycle, which I think is good for now
		Path:   "/",
	})
	if err != nil {
		return fmt.Errorf("creating cookie session for oidc middleware: %w", err)
	}

	mux := http.NewServeMux()

	mux.Handle("/keys", keysh)
	mux.Handle("/.well-known/openid-configuration", discoh)

	heh := &httpErrHandler{}

	// TODO - usernameless via resident keys would be nice, but need to
	// see what support is like.
	rrk := false
	wn, err := webauthn.New(&webauthn.Config{
		RPDisplayName: issuer.URL.Hostname(), // Display Name for your site
		RPID:          issuer.URL.Hostname(), // Generally the FQDN for your site
		RPOrigin:      issuer.URL.String(),   // The origin URL for WebAuthn requests
		AuthenticatorSelection: protocol.AuthenticatorSelection{
			UserVerification:   protocol.VerificationRequired,
			RequireResidentKey: &rrk,
		},
	})
	if err != nil {
		return fmt.Errorf("configuring webauthn: %w", err)
	}

	// start configuration of webauthn manager
	oidcSess, err := oidcm.NewMemorySessionStore(http.Cookie{Name: "oidc-login", Path: "/", HttpOnly: true})
	if err != nil {
		return fmt.Errorf("creating OIDC middleware session store: %w", err)
	}
	mgr := &webauthnManager{
		db:       db,
		webauthn: wn,
		sessmgr:  webSessMgr,
		oidcMiddleware: &oidcm.Handler{
			Issuer:       issuer.URL.String(),
			ClientID:     uuid.New().String(), // TODO - something that will live beyond restarts
			ClientSecret: uuid.New().String(),
			BaseURL:      issuer.URL.String(),
			RedirectURL:  issuer.URL.String() + "/local-oidc-callback",
			SessionStore: oidcSess,
		},
		csrfMiddleware: nosurf.NewPure,
		// admins: p.Webauthn.AdminSubjects, // TODO - google account id
		acrs: nil,
	}
	// this is a dumb hack, because we use the middleware super
	// restrictively but it needs to catch it's callback.
	mux.Handle("/local-oidc-callback", mgr.oidcMiddleware.Wrap(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("should never get here?"))
	})))

	clients.sources = append([]core.ClientSource{
		&staticClients{
			clients: []clientConfig{
				{
					ClientID:     mgr.oidcMiddleware.ClientID,
					ClientSecret: []string{mgr.oidcMiddleware.ClientSecret},
					RedirectURL:  []string{mgr.oidcMiddleware.RedirectURL},
				},
			},
		},
	}, clients.sources...)

	mgr.AddHandlers(mux)

	svr := oidcServer{
		issuer:          issuer.URL.String(),
		oidcsvr:         oidcsvr,
		eh:              heh,
		tokenValidFor:   15 * time.Minute,
		refreshValidFor: 12 * time.Hour,
		sessmgr:         webSessMgr,
		// upstreamPolicy:  []byte(ucp),
		webauthn: wn,
		db:       db,
	}

	pubContent, err := fs.Sub(fs.FS(staticFiles), "web")
	if err != nil {
		return fmt.Errorf("creating public subfs: %w", err)
	}
	fs := http.FileServer(http.FS(pubContent))
	mux.Handle("/public/", fs)

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("OK"))
	})

	_, loopback, err := net.ParseCIDR("127.0.0.0/8")
	if err != nil {
		return err
	}
	mux.HandleFunc("/reloaddb", func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !loopback.Contains(net.ParseIP(ip)) {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		if err := db.Reload(); err != nil {
			slog.ErrorContext(r.Context(), "database reload failed", slog.Any("error", err))
			http.Error(w, fmt.Sprintf("reload failed: %v", err), http.StatusInternalServerError)
		} else {
			slog.InfoContext(r.Context(), "database reloaded")
		}
	})

	svr.AddHandlers(mux)

	var g run.Group

	g.Add(run.SignalHandler(ctx, os.Interrupt, unix.SIGTERM))

	g.Add(ksm.Run, ksm.Interrupt)

	// this will always try and create a session for discovery and stuff,
	// but we shouldn't save it. but, we need it for logging and stuff. TODO
	// at some point consider splitting the middleware, but then we might
	// need to dup the middleware wrap or something.
	hh := baseMiddleware(mux, webSessMgr, oidcmSessMgr)

	hs := &http.Server{
		Addr:    addr,
		Handler: hh,
	}

	g.Add(func() error {
		slog.Info("server listing", slog.String("addr", "http://"+addr))
		if err := hs.ListenAndServe(); err != nil {
			return fmt.Errorf("serving http: %v", err)
		}
		return nil
	}, func(error) {
		// new context for this, parent is likely already shut down
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		_ = hs.Shutdown(ctx)
	})

	return g.Run()
}

func registrationURL(iss *url.URL, user User) *url.URL {
	u := *iss
	if !strings.HasSuffix(u.Path, "/") {
		u.Path += "/"
	}
	u2, err := u.Parse("/registration")
	if err != nil {
		panic(err)
	}
	q := u2.Query()
	q.Add("user_id", user.ID)
	q.Add("enrollment_token", user.EnrollmentKey)
	u2.RawQuery = q.Encode()
	return u2
}

// activateUser marks the user as activated and deletes its enrollment key.
// Must be called after the user has completed the registration flow.
func activateUser(db *DB, userID string) error {
	user, err := db.GetUserByID(userID)
	if err != nil {
		return fmt.Errorf("get user %s: %w", userID, err)
	}

	user.EnrollmentKey = ""
	user.Activated = true

	if err := db.UpdateUser(user); err != nil {
		return fmt.Errorf("update user %s: %w", userID, err)
	}

	fmt.Println("Done.")

	return nil
}

// reloadDB tells the server running on addr to reload its database from disk.
func reloadDB(addr string) {
	resp, err := http.Get("http://" + addr + "/reloaddb")
	if err != nil {
		fatalf("database reload failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		fatalf("database reload failed: %s", string(b))
	}
}

func fatal(s string) {
	fmt.Fprintf(os.Stderr, "webauthn-oidc-idp: %s\n", s)
	os.Exit(1)
}

func fatalf(s string, args ...any) {
	fmt.Fprintf(os.Stderr, fmt.Sprintf("webauthn-oidc-idp: %s\n", s), args...)
	os.Exit(1)
}

func logErr(err error) slog.Attr {
	return slog.Any("error", err)
}
