package constants

var MapOfListeners map[string]bool = map[string]bool{
	"nats":   true,
	"awssqs": false,
	"kafka":  false,
}

const (
	GmailClientId       = "1031026745675-fq5hu7unat9omdi26h143io6dkit44m2.apps.googleusercontent.com"
	GmailClientSecret   = "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"
	OutlookClientId     = "acc58ddf-1662-46de-8f67-0a6d7fbb5f4e"
	OutlookClientSecret = "0Xh8Q~CawadUSmvZCE_..gpig3sdtENm4CEkTdmj"
	RedirectURI         = "http://localhost:8080/oauth2callback"
)
