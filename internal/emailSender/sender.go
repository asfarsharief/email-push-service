package emailsender

type SenderInterface interface {
	SendMail(fromAddress, toAddress, subject, body string) error
	FetchAccessToken() string
}

func GetEmailSenderObject(provider string) SenderInterface {
	switch provider {
	case "gmail":
		return NewGmailSender()
	}
	return nil
}
