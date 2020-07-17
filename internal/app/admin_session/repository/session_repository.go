package repository

import (
	"time"

	"github.com/go-redis/redis"
)

type SessionStorage struct {
	sessionDB      *redis.Client
	expirationTime time.Duration
}

func (ss *SessionStorage) AddCookie(cookieValue, adminUsername string) error {
	return ss.sessionDB.Set(cookieValue, adminUsername, ss.expirationTime).Err()
}

func (ss *SessionStorage) GetUsernameByCookie(cookieValue string) (string, error) {
	return ss.sessionDB.Get(cookieValue).Result()
}

func (ss *SessionStorage) DeleteCookie(cookieValue string) error {
	return ss.sessionDB.Del(cookieValue).Err()
}
