package store

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenInfo struct {
	ClientID uuid.UUID `json:"client_id"`
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	Nickname string    `json:"nickname"`
	Role     string    `json:"role"`
}

type Claims struct {
	TokenInfo
	jwt.RegisteredClaims
}
