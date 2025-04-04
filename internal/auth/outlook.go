package auth

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"

// 	"golang.org/x/oauth2"
// )

// var clientID = "acc58ddf-1662-46de-8f67-0a6d7fbb5f4e"
// var clientSecret = "0Xh8Q~CawadUSmvZCE_..gpig3sdtENm4CEkTdmj"
// var redirectURI = "http://localhost:8080/callback-outlook"

// var oauthConfig = oauth2.Config{
// 	ClientID:     clientID,
// 	ClientSecret: clientSecret,
// 	RedirectURL:  redirectURI,
// 	Scopes:       []string{"Mail.Send", "offline_access", "User.Read"},
// 	Endpoint: oauth2.Endpoint{
// 		AuthURL:  "https://login.microsoftonline.com/common/oauth2/v2.0/authorize",
// 		TokenURL: "https://login.microsoftonline.com/common/oauth2/v2.0/token",
// 	},
// }

// // curl -X POST https://login.microsoftonline.com/common/oauth2/v2.0/token?client_id=acc58ddf-1662-46de-8f67-0a6d7fbb5f4e&client_secret=0Xh8Q~CawadUSmvZCE_..gpig3sdtENm4CEkTdmj&code=708303&redirect_uri=http://localhost:8080/callback-outlook&grant_type=authorization_code"

// // https://login.microsoftonline.com/organizations/oauth2/v2.0/authorize?client_id=acc58ddf-1662-46de-8f67-0a6d7fbb5f4e&response_type=code&redirect_uri=http://localhost:8080/callback-outlook&scope=Mail.Send offline_access User.Read&response_mode=query

// func main() {
// 	authURL := oauthConfig.AuthCodeURL("random_state_string", oauth2.AccessTypeOffline)
// 	fmt.Println("🔗 Open this URL in your browser:\n", authURL)

// 	http.HandleFunc("/callback-outlook", handleCallback)
// 	log.Println("🚀 Waiting for authentication...")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// func handleCallback(w http.ResponseWriter, r *http.Request) {
// 	code := r.URL.Query().Get("code")
// 	fmt.Println("Authorization Code:", code)
// 	if code == "" {
// 		http.Error(w, "No code in URL", http.StatusBadRequest)
// 		return
// 	}

// 	token, err := oauthConfig.Exchange(context.Background(), code)
// 	if err != nil {
// 		http.Error(w, "Token exchange failed: "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Save token to file
// 	file, _ := os.Create("token.json")
// 	defer file.Close()
// 	json.NewEncoder(file).Encode(token)

// 	fmt.Fprintf(w, "✅ Authentication successful! Token saved.")
// 	log.Println("✅ Token saved to token.json")
// }
