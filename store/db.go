package store

import (
	"database/sql"
	"email-push-service/pkg/logger"
	"fmt"

	_ "modernc.org/sqlite"
)

type DbStore struct {
	Conn *sql.DB
}

type DbStoreInterface interface {
	GetUsers(userId, tenantId string) (*Users, error)
	FetchQuotaByTenant(tenantId string) (*QuotaTracking, error)
	InsertOrUpdateQuotaTracking(quotaTracking *QuotaTracking) error
}

func NewSqliteDbStore() DbStoreInterface {
	db, err := sql.Open("sqlite", "./store/database.db")
	if err != nil {
		logger.Error("failed to open database: %w", err)
		return nil
	}

	// Create a sample table if not exists
	logger.Info("Database initialized successfully")
	return &DbStore{Conn: db}
}

func (ds *DbStore) GetUsers(userId, tenantId string) (*Users, error) {
	rows, err := ds.Conn.Query("SELECT userId, tenantId, email FROM Users where userId = ? and tenantId = ?", userId, tenantId)
	if err != nil {
		logger.Errorf("failed to fetch users: %w", err)
		return nil, err
	}
	defer rows.Close()

	user := Users{}
	for rows.Next() {
		if err := rows.Scan(&user.UserId, &user.TenantId, &user.EmailId); err != nil {
			fmt.Println("error: ", err)
			return nil, err
		}

		break
	}
	return &user, nil
}

func (ds *DbStore) FetchQuotaByTenant(tenantId string) (*QuotaTracking, error) {
	rows, err := ds.Conn.Query("SELECT tenantId, date, emailsSent, dailyLimit, quotaMultiplier FROM QuotaTracking where tenantId = ? ", tenantId)
	if err != nil {
		logger.Errorf("failed to fetch users: %w", err)
		return nil, err
	}
	defer rows.Close()

	quotaTracking := QuotaTracking{}
	for rows.Next() {
		if err := rows.Scan(&quotaTracking.TenantId, &quotaTracking.Date, &quotaTracking.EmailsSent, &quotaTracking.DailyLimit, &quotaTracking.QuotaMultiplier); err != nil {
			fmt.Println("error: ", err)
			return nil, err
		}

		break
	}
	return &quotaTracking, nil
}

func (ds *DbStore) InsertOrUpdateQuotaTracking(quotaTracking *QuotaTracking) error {
	stmt := `
INSERT INTO QuotaTracking (tenantId, date, emailsSent, dailyLimit, quotaMultiplier)
VALUES (?, ?, ?, ?, ?)
ON CONFLICT(tenantId, date)
DO UPDATE SET
    emailsSent = excluded.emailsSent,
    dailyLimit = excluded.dailyLimit,
	quotaMultiplier = excluded.quotaMultiplier;
`

	_, err := ds.Conn.Exec(stmt, quotaTracking.TenantId, quotaTracking.Date, quotaTracking.EmailsSent, quotaTracking.DailyLimit, quotaTracking.QuotaMultiplier)
	if err != nil {
		logger.Error("Upsert failed:", err)
	}
	return err
}
