package server

import (
	"email-push-service/internal/auth"
	"email-push-service/pkg/logger"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome! <a href='/login'>Login with Provider(gmail or outlook)</a>")
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	provider := r.URL.Query().Get("provider")
	if provider == "" {
		http.Error(w, "provider not found in URL", http.StatusBadRequest)
		return
	}
	authObject := auth.GetAuthStruct(provider)
	if authObject == nil {
		http.Error(w, "invalid provider found in URL", http.StatusBadRequest)
		return
	}
	url := authObject.GetLoginUrl()
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code not found in URL", http.StatusBadRequest)
		return
	}

	authObject := auth.GetAuthStruct("gmail")
	if authObject == nil {
		http.Error(w, "invalid provider found in URL", http.StatusBadRequest)
		return
	}
	// Exchange the authorization code for access and refresh tokens
	token, err := authObject.HandleCallback(code)
	if err != nil {
		http.Error(w, "Token exchange failed", http.StatusInternalServerError)
		return
	}

	// Save token to file (or store in DB)
	saveToken(token)

	// Display token info
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

// Save token to file (you can use a database instead)
func saveToken(token *oauth2.Token) {
	file, err := os.Create("token.json")
	if err != nil {
		logger.Error("Unable to create token file:", err)
		return
	}
	defer file.Close()

	json.NewEncoder(file).Encode(token)
	logger.Info("Token saved successfully!")
}
