package emailprocessor

import (
	emailsender "email-push-service/internal/emailSender"
	"email-push-service/pkg/logger"
	"email-push-service/store"
	"encoding/json"
	"net/mail"
	"time"
)

var db store.DbStoreInterface

func init() {
	db = store.NewSqliteDbStore()
}

type SendEmailStruct struct {
	ToAddress string `json:"toAddress"`
	TenantId  string `json:"tenantId"`
	UserId    string `json:"userId"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
}

func ProcessRequest(message []byte) {
	logger.Info("Starting processing...")
	emailData := new(SendEmailStruct)

	err := json.Unmarshal(message, emailData)
	if err != nil {
		logger.Error(err)
		return
	}

	user, err := db.GetUsers(emailData.UserId, emailData.TenantId)
	if err != nil {
		logger.Error("Error in GetUsers: ", err)
		return
	}

	quota, err := db.FetchQuotaByTenant(emailData.TenantId)
	if err != nil {
		logger.Error("Error in FetchQuotaByTenant: ", err)
		return
	}

	now := time.Now()
	if quota == nil {
		quota = &store.QuotaTracking{
			TenantId:   emailData.TenantId,
			Date:       now.Format("2006-01-02"),
			EmailsSent: 0,
			DailyLimit: 50,
		}
	} else {
		t, err := time.Parse("2006-01-02", quota.Date)
		if err != nil {
			quota = &store.QuotaTracking{
				TenantId:        emailData.TenantId,
				Date:            now.Format("2006-01-02"),
				QuotaMultiplier: 2,
				EmailsSent:      0,
				DailyLimit:      50,
			}
		} else {
			if !isWithinOneDay(t, now) {
				quota.EmailsSent = 0
				quota.Date = now.Format("2006-01-02")
				if quota.DailyLimit < 100 {
					quota.DailyLimit *= quota.QuotaMultiplier
					quota.QuotaMultiplier *= quota.QuotaMultiplier
				}
			}
		}
	}

	logger.Infof("Quota for tenant: %v", quota)
	if quota.EmailsSent > quota.DailyLimit {
		//Push back of the queue. Nats doesn't have feature, can be done for others
		return
	}

	//validate
	_, err = mail.ParseAddress(emailData.ToAddress)
	if err != nil {
		logger.Error("Error in ParseAddress: ", emailData.ToAddress, ", error: ", err)
		return
	}

	senderObject := emailsender.GetEmailSenderObject("gmail")
	if senderObject == nil {
		logger.Error("Error in creating MailSenderObject")
		return
	}

	err = senderObject.SendMail(user.EmailId, emailData.ToAddress, emailData.Subject, emailData.Body)
	if err != nil {
		logger.Error("Error while sending mail: ", err)
		//push to the back of the queue
	} else {
		quota.EmailsSent++
	}

	db.InsertOrUpdateQuotaTracking(quota)
	logger.Info("Processed request successfully")
}

func isWithinOneDay(a, b time.Time) bool {
	diff := a.Sub(b)
	if diff < 0 {
		diff = -diff // take absolute value
	}
	return diff <= 24*time.Hour
}
