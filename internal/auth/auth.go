package auth

import (
	"context"
	"email-push-service/pkg/constants"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthStruct struct {
	oauthConfig *oauth2.Config
}

type AuthInterface interface {
	GetLoginUrl() string
	HandleCallback(code string) (*oauth2.Token, error)
}

func GetAuthStruct(provider string) AuthInterface {
	switch provider {
	case "gmail":
		return &AuthStruct{
			oauthConfig: &oauth2.Config{
				ClientID:     constants.GmailClientId,
				ClientSecret: constants.GmailClientSecret,
				RedirectURL:  constants.RedirectURI,
				Scopes:       []string{"https://www.googleapis.com/auth/gmail.send"},
				Endpoint:     google.Endpoint,
			},
		}
	case "outlook":
		return &AuthStruct{
			oauthConfig: &oauth2.Config{
				ClientID:     constants.OutlookClientId,
				ClientSecret: constants.GmailClientSecret,
				RedirectURL:  constants.RedirectURI,
				Scopes:       []string{"Mail.Send", "offline_access", "User.Read"},
				Endpoint: oauth2.Endpoint{
					AuthURL:  "https://login.microsoftonline.com/common/oauth2/v2.0/authorize",
					TokenURL: "https://login.microsoftonline.com/common/oauth2/v2.0/token",
				},
			},
		}
	}
	return nil
}

func (as *AuthStruct) GetLoginUrl() string {
	return as.oauthConfig.AuthCodeURL("random_state_string", oauth2.AccessTypeOffline)
}

func (as *AuthStruct) HandleCallback(code string) (*oauth2.Token, error) {
	// Exchange the authorization code for access and refresh tokens
	token, err := as.oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// Save token to file (you can use a database instead)
// func saveToken(token *oauth2.Token) {
// 	file, err := os.Create("token.json")
// 	if err != nil {
// 		log.Fatal("Unable to create token file:", err)
// 	}
// 	defer file.Close()

// 	json.NewEncoder(file).Encode(token)
// 	fmt.Println("Token saved successfully!")
// }
