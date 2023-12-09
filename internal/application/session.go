package application

import (
	"encoding/json"
	"errors"

	"github.com/buemura/temp-storage/internal/domain/cache"
	"github.com/buemura/temp-storage/internal/domain/session"
)

type SessionService struct {
	cacheStorage cache.CacheStorage
}

func NewSessionService(cacheStorage cache.CacheStorage) *SessionService {
	return &SessionService{
		cacheStorage: cacheStorage,
	}
}

func (ss *SessionService) GetSession(sessionId string) (*session.Session, error) {
	val := ss.cacheStorage.Get("sessionId:" + sessionId)
	if val == "" {
		return nil, errors.New(session.SessionNotFoundError)
	}

	var sess *session.Session
	json.Unmarshal([]byte(val), &sess)

	return sess, nil
}

func (ss *SessionService) CreateSession() (*session.Session, error) {
	sess := session.NewSession(10, 10)
	jsonBytes, err := json.Marshal(sess)
	if err != nil {
		return nil, err
	}
	sessStr := string(jsonBytes)

	err = ss.cacheStorage.Set("sessionId:"+sess.ID, sessStr, sess.TimeToLive)
	if err != nil {
		return nil, err
	}
	return sess, nil
}

func (ss *SessionService) UpdateSession(sess *session.Session) (*session.Session, error) {
	jsonBytes, err := json.Marshal(sess)
	if err != nil {
		return nil, err
	}
	sessStr := string(jsonBytes)

	err = ss.cacheStorage.Set("sessionId:"+sess.ID, sessStr, sess.TimeToLive)
	if err != nil {
		return nil, err
	}

	return sess, nil
}
