package store

import (
	"github.com/OpsOMI/S.L.A.M/internal/shared/store"
	"github.com/golang-jwt/jwt/v5"
)

type SessionStore struct {
	JWT  string
	Room string
	store.TokenInfo
}

func NewSessionStore() *SessionStore {
	return &SessionStore{}
}

func (s *SessionStore) SetToken(token string) {
	s.JWT = token
}

func (s *SessionStore) SetRoom(room string) {
	s.Room = room
}

func (s *SessionStore) GetRoom() string {
	return s.Room
}

func (s *SessionStore) GetToken() string {
	return s.JWT
}

func (s *SessionStore) ParseJWT() error {
	if s.JWT == "" {
		return nil
	}

	claims := &store.Claims{}

	_, _, err := new(jwt.Parser).ParseUnverified(s.JWT, claims)
	if err != nil {
		return err
	}

	s.TokenInfo = claims.TokenInfo

	return nil
}
