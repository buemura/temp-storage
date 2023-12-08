package session

import (
	"github.com/google/uuid"
)

type Session struct {
	ID         string `json:"id"`
	TimeToLive int    `json:"timeToLive"`
	MaxSize    int    `json:"maxSize"`
}

func NewSession(ttl, maxSize int) *Session {
	return &Session{
		ID:         uuid.NewString(),
		TimeToLive: ttl,
		MaxSize:    maxSize,
	}
}
