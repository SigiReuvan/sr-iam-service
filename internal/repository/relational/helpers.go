package relational

import (
	"fmt"

	"github.com/SigiReuvan/iam-service/internal"
)

func (repo *repository) checkUniqueness(key string, value string) (string, error) {
	var users []internal.User
	query := fmt.Sprintf("%s = ?", key)
	result := repo.db.Table("Users").Where(query, value).Find(&users)
	if result.Error != nil {
		return "failed", result.Error
	}
	if len(users) > 0 {
		return "failed", nil
	}
	return "success", nil
}
