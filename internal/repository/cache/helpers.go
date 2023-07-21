package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/SigiReuvan/iam-service/internal"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func (repo *repository) createSession(ctx context.Context, u internal.User) {
	id := uuid.New().String()

	creationTime := time.Now()
	expirationTime := time.Now().Add(30 * time.Minute)

	s := internal.ActiveSession{
		UID:            u.ID,
		Role:           u.Role,
		Creation:       creationTime,
		Expiration:     expirationTime,
		SessionActions: []internal.SessionAction{},
	}

	repo.cache.Set(ctx, id, s, 30*time.Minute)
}

func (repo *repository) addAction(ctx context.Context, u internal.User, action string, resource string) error {
	val, err := repo.cache.Get(ctx, u.ID).Result()
	if err == redis.Nil {
		return redis.Nil
	} else if err != nil {
		return err
	}

	var session *internal.ActiveSession
	err = json.Unmarshal([]byte(val), session)
	if err != nil {
		return err
	}

	session = &internal.ActiveSession{
		Expiration: time.Now().Add(30 * time.Minute),
	}

	session.SessionActions = append(session.SessionActions, internal.SessionAction{
		Action:       action,
		Resource:     resource,
		LastActivity: time.Now(),
	})

	repo.cache.Set(ctx, u.ID, session, 30*time.Minute)

	return nil
}
