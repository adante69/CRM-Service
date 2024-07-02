package handlers

import (
	"CRM-Service/config"
	"CRM-Service/internal/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func RegisterAuthRoutes(router *mux.Router, authHandler *AuthHandler) {
	router.HandleFunc("/auth/signup", authHandler.SignUp).Methods("POST")
	router.HandleFunc("/auth/signin", authHandler.SignIn).Methods("POST")
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.AuthService.Register(creds.Email, creds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	conf, err := config.LoadConfiguration()

	token, err := h.AuthService.Login(conf, creds.Email, creds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(24 * time.Hour),
	})
}
