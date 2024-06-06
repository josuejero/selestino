// api/middleware_test.go

package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestRoleBasedAuthorization(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	adminToken, _ := GenerateToken("admin_user", "admin")
	req, _ := http.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{Name: "token", Value: adminToken})

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/", RoleBasedAuthorization("admin")(handler))
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	userToken, _ := GenerateToken("regular_user", "user")
	req, _ = http.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{Name: "token", Value: userToken})

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusForbidden, rr.Code)
}
