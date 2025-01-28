package middleware

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

const (
	AuthName  = "Authorization"
	AuthLogin = "X-User-Agent"
	TokenExp  = time.Hour * 24
)

func JwtAuthMiddleware(secret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

			c, err := r.Cookie(AuthName)
			if err != nil {
				if r.Method == http.MethodGet {
					log.Error("is unauthorized", slog.Any("error", err))
					w.WriteHeader(http.StatusUnauthorized)
					next.ServeHTTP(w, r)
					return
				} else {
					writeNewCookie(w, r, next, secret)
					log.Debug("new cookie", slog.Any("error", err))
					return
				}
			}

			t := c.Value
			valid, err := isAuthorized(t, secret)
			if !valid {
				writeNewCookie(w, r, next, secret)
				log.Debug("is authorized", slog.Any("error", err))
				return
			}

			mustbeuser, err := mustUserID(t, secret)

			if !mustbeuser {
				w.WriteHeader(http.StatusUnauthorized)
				next.ServeHTTP(w, r)
				log.Error("is authorized", slog.Any("error", err))
				return
			}

			userID, _ := getUserID(t, secret)
			r.Header.Set(AuthLogin, userID)

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

func writeNewCookie(w http.ResponseWriter, r *http.Request, next http.Handler, secret string) {
	user := NewUser()

	token, err := createAccessToken(user.UserID, secret, TokenExp)
	if err != nil {
		slog.Error("create access token", slog.Any("error", err))
	}

	cookie := http.Cookie{
		Name:  AuthName,
		Value: token,
	}
	r.Header.Set(AuthLogin, user.UserID)
	http.SetCookie(w, &cookie)
	next.ServeHTTP(w, r)
}

type User struct {
	UserID string
}

func NewUser() *User {
	return &User{
		UserID: uuid.New().String(),
	}
}
