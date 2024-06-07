// api/router.go

package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/josuejero/selestino/internal/models"
	"github.com/josuejero/selestino/internal/repository"
)

var userRepo *repository.UserRepository

func InitializeRouter(db *sql.DB) *mux.Router {
	userRepo = &repository.UserRepository{DB: db}
	SetRecipeRepo(&repository.RecipeRepository{DB: db})

	router := mux.NewRouter()

	// Define rate limiter
	rateLimiter := NewRateLimiter(1000, 10, time.Minute)

	// Define your API routes here
	router.HandleFunc("/recipes", GetRecipes).Methods("GET")
	router.HandleFunc("/recipes", AddRecipe).Methods("POST")
	router.HandleFunc("/recipes/search", SearchRecipesByCriteria).Methods("GET")

	router.HandleFunc("/register", RegisterUser).Methods("POST")
	router.HandleFunc("/login", LoginUser).Methods("POST")

	// Apply rate limiter middleware to all routes
	router.Use(rateLimiter.Limit)

	// Apply role-based middleware to protected routes
	adminRoutes := router.PathPrefix("/admin").Subrouter()
	adminRoutes.Use(RoleBasedAuthorization("admin"))
	adminRoutes.HandleFunc("/recipes", AddRecipe).Methods("POST")

	return router
}

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := userRepo.CreateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authenticatedUser, authenticated, err := userRepo.AuthenticateUser(user.Username, user.Password)
	if err != nil || !authenticated {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: authenticatedUser.Username,
		Role:     authenticatedUser.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
