package store

type Users struct {
	UserId   string
	TenantId string
	EmailId  string
}

type CredStore struct {
	TenantId     string
	Provider     string
	AccessToken  string
	RefreshToken string
	ExpiresAt    string
}

type QuotaTracking struct {
	TenantId        string
	Date            string
	EmailsSent      int
	DailyLimit      int
	QuotaMultiplier int
}

type AuthCredentialStore struct {
	TenantId     string
	Provider     string
	ClientId     string
	ClientSecret string
}

var MapOfUsers map[[2]string]Users = map[[2]string]Users{
	{"asfar.sharief", "company1"}: {
		UserId:   "asfar.sharief",
		TenantId: "company1",
		EmailId:  "asfarsharief015@gmail.com",
	},
}

func FetchUser(userName, tenantId string) *Users {
	if user, ok := MapOfUsers[[2]string{userName, tenantId}]; ok {
		return &user
	}
	return nil
}

var MapOfCreds map[string]CredStore = map[string]CredStore{
	"company1": {
		TenantId: "company1",
		Provider: "gmail",
	},
}

var MapOfAuthCreds map[string]AuthCredentialStore = map[string]AuthCredentialStore{
	"company1": {
		TenantId:     "company1",
		Provider:     "gmail",
		ClientId:     "",
		ClientSecret: "",
	},
}

var MapOfQuota map[string]QuotaTracking = map[string]QuotaTracking{
	"company1": {
		TenantId:   "company1",
		EmailsSent: 0,
		DailyLimit: 50,
	},
}

func FetchQuotaByTenant(tenantId string) *QuotaTracking {
	if data, ok := MapOfQuota[tenantId]; ok {
		return &data
	}
	return nil
}
