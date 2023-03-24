package session

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
)

type Repository struct {
	mu        sync.RWMutex
	storage   map[uuid.UUID]domain.Session
	redisConn redis.Conn
}

func NewRepository(redisConn redis.Conn) Repository {
	return Repository{storage: map[uuid.UUID]domain.Session{}, redisConn: redisConn}
}

func (repo *Repository) Add(session domain.Session) error {
	if time.Time(session.ExpiresAt).Before(time.Now()) {
		return nil
	}
	dto := sessionDTO{
		ExpiresAtString: session.ExpiresAt.Format(time.RFC3339),
		UserID:          session.UserID,
	}

	// TODO: сериализую только user_id
	// 1. в session еще есть поля,
	// 2. иммет ли смысл сериализовать UserAgent? (в лекциях его сериализовать)

	dataSerialized, err := json.Marshal(map[string]interface{}{
		"user_id":    dto.UserID,
		"expires_at": dto.ExpiresAtString,
	})
	if err != nil {
		return err
	}
	// 86400s = 24h
	// TODO: session.ID или session.ID.String()
	result, err := redis.String(repo.redisConn.Do("SET", session.ID, dataSerialized, "EX", 86400))
	if err != nil {
		return fmt.Errorf("cant set data in redis:", err)
	}
	if result != "OK" {
		return fmt.Errorf("'set' in redis replies 'not OK'")
	}
	return nil
}

func (repo *Repository) Get(sessionID uuid.UUID) (domain.Session, error) {
	dto := sessionDTO{}

	data, err := redis.Bytes(repo.redisConn.Do("GET", sessionID))
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			return domain.Session{}, domain.ErrSessionNotFound
		}
		return domain.Session{}, fmt.Errorf("cant get data in redis: %w", err)
	}
	err = json.Unmarshal(data, &dto)
	if err != nil {
		return domain.Session{}, fmt.Errorf("cant unpack session data from redis: %w", err)
	}
	time, err := time.Parse(time.RFC3339, dto.ExpiresAtString)
	if err != nil {
		return domain.Session{}, nil
	}
	session := domain.Session{
		ExpiresAt: time,
		UserID:    dto.UserID,
		ID:        sessionID,
	}
	return session, nil
}

func (repo *Repository) Delete(sessionID uuid.UUID) error {
	// TODO: может можно лучше обработать ошибку, зачем приводить к Int? result != OK?
	_, err := redis.Int(repo.redisConn.Do("DEL", sessionID))
	if err != nil {
		return fmt.Errorf("cant delete by redis: %w", err)
	}
	return nil
}
