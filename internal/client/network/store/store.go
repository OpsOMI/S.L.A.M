package store

import (
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/tokenstore"
	"github.com/golang-jwt/jwt/v5"
)

type SessionStore struct {
	JWT string
	tokenstore.TokenInfo
}

func NewSessionStore() *SessionStore {
	return &SessionStore{}
}

func (s *SessionStore) SetToken(token string) {
	s.JWT = token
}

func (s *SessionStore) GetToken() string {
	return s.JWT
}

func (s *SessionStore) ParseJWT() error {
	if s.JWT == "" {
		return nil
	}

	claims := &tokenstore.Claims{}

	_, _, err := new(jwt.Parser).ParseUnverified(s.JWT, claims)
	if err != nil {
		return err
	}

	s.TokenInfo = claims.TokenInfo

	return nil
}
