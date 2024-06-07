// api/middleware_test.go

package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/josuejero/selestino/api"
	"github.com/stretchr/testify/assert"
)

func TestRoleBasedAuthorization(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	adminToken, _ := api.GenerateToken("admin_user", "admin")
	req, _ := http.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{Name: "token", Value: adminToken})

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/", api.RoleBasedAuthorization("admin")(handler))
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	userToken, _ := api.GenerateToken("regular_user", "user")
	req, _ = http.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{Name: "token", Value: userToken})

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusForbidden, rr.Code)
}
