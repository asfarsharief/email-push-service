package emailsender

import (
	"context"
	"email-push-service/pkg/logger"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type GmailSender struct {
}

func NewGmailSender() *GmailSender {
	return &GmailSender{}
}

func (gs *GmailSender) FetchAccessToken() string {
	// Check if file exists
	if _, err := os.Stat("./token.json"); os.IsNotExist(err) {
		fmt.Println("token.json does not exist")
		return ""
	}

	// Read file
	data, err := os.ReadFile("./token.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}

	// Parse JSON
	var tokenData map[string]interface{}
	if err := json.Unmarshal(data, &tokenData); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return ""
	}

	// Extract access_token
	accessToken, ok := tokenData["access_token"].(string)
	if !ok {
		fmt.Println("access_token not found or invalid type")
		return ""
	}
	return accessToken
}

func (gs *GmailSender) SendMail(fromAddress, toAddress, subject, body string) error {
	accessToken := gs.FetchAccessToken()
	fmt.Println("Access: ", accessToken)
	if accessToken == "" {
		logger.Info("Token file missing, call localhost:8080/login?provider=gmail to download file")
		return errors.New("Token file missing")
	}
	token := &oauth2.Token{AccessToken: accessToken}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(token)
	client := oauth2.NewClient(ctx, ts)

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		logger.Errorf("Unable to create Gmail client: %v", err)
		return err
	}

	email := fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n\r\n%s", fromAddress, toAddress, subject, body)

	var message gmail.Message
	message.Raw = base64.URLEncoding.EncodeToString([]byte(email))

	_, err = srv.Users.Messages.Send(fromAddress, &message).Do()
	if err != nil {
		logger.Errorf("Unable to send email: %v", err)
		return err
	}

	logger.Info("Email sent successfully!")
	return nil
}
