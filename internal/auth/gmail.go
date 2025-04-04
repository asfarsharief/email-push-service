package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthConfig *oauth2.Config
)

// Load Google OAuth configuration
func init() {
	oauthConfig = &oauth2.Config{
		ClientID:     "1031026745675-fq5hu7unat9omdi26h143io6dkit44m2.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8080/oauth2callback",
		Scopes:       []string{"https://www.googleapis.com/auth/gmail.send"},
		Endpoint:     google.Endpoint,
	}
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome! <a href='/login'>Login with Google</a>")
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	url := oauthConfig.AuthCodeURL("random_state_string", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code not found in URL", http.StatusBadRequest)
		return
	}

	// Exchange the authorization code for access and refresh tokens
	token, err := oauthConfig.Exchange(context.Background(), code)
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
		log.Fatal("Unable to create token file:", err)
	}
	defer file.Close()

	json.NewEncoder(file).Encode(token)
	fmt.Println("Token saved successfully!")
}
