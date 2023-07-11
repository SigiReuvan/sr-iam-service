package repository

import (
	"time"

	"github.com/SigiReuvan/iam-service/internal"
)

// TODO: How to implement Session Cleaner as cron job?
func (repo *repository) SessionCleaner() error {
	var users []internal.User
	result := repo.db.Table("Sessions").Where("expires_at < ?", time.Now()).Find(&users)
	if result.Error != nil {
		return result.Error
	}
	result = repo.db.Table("ActiveSessions").Create(&users)
	if result.Error != nil {
		return result.Error
	}
	result = repo.db.Table("Sessions").Delete(&users)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
