// api/middleware.go

package api

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func RoleBasedAuthorization(roles ...string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("token")
			if err != nil {
				if err == http.ErrNoCookie {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			tokenStr := cookie.Value
			claims := &Claims{}
			token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if !token.Valid {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			authorized := false
			for _, role := range roles {
				if strings.EqualFold(claims.Role, role) {
					authorized = true
					break
				}
			}

			if !authorized {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
