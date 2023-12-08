package session

import (
	"github.com/buemura/temp-storage/internal/domain/file"
	"github.com/google/uuid"
)

type Session struct {
	ID         string       `json:"id"`
	TimeToLive int          `json:"timeToLive"`
	MaxSize    int          `json:"maxSize"`
	Files      []*file.File `json:"files"`
}

func NewSession(ttl, maxSize int) *Session {
	return &Session{
		ID:         uuid.NewString(),
		TimeToLive: ttl,
		MaxSize:    maxSize,
	}
}

func (s *Session) AddFile(f *file.File) {
	s.Files = append(s.Files, f)
}
