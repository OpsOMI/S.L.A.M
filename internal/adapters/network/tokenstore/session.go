package tokenstore

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenInfo struct {
	ClientID string `json:"client_id"`
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}

type Claims struct {
	TokenInfo
	jwt.RegisteredClaims
}

type ITokenStore interface {
	GenerateToken(
		clientID, userID, username, nickname string,
		duration time.Duration,
	) (string, error)

	ValidateToken(
		tokenStr *string,
	) (*Claims, error)

	ParseToken(
		tokenStr *string,
	) *TokenInfo
}

type manager struct {
	issuer string
	secret []byte
}

func NewJWTManager(
	issuer, secret string,
) ITokenStore {
	return &manager{
		issuer: issuer,
		secret: []byte(secret),
	}
}

func (m *manager) GenerateToken(clientID, userID, username, nickname string, duration time.Duration) (string, error) {
	claims := Claims{
		TokenInfo: TokenInfo{
			ClientID: clientID,
			UserID:   userID,
			Username: username,
			Nickname: nickname,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    m.issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(m.secret)
}

func (m *manager) ValidateToken(
	tokenStr *string,
) (*Claims, error) {
	if tokenStr == nil {
		return nil, errors.New("jwt.invalid_token")
	}

	token, err := jwt.ParseWithClaims(*tokenStr, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("jwt.unexpected_signing_method")
		}
		return m.secret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("jwt.invalid_token")
	}
	return claims, nil
}

// This function assumes the JWT is already validated by middleware, so it skips error handling and returns the parsed token info directly.
func (m *manager) ParseToken(tokenStr *string) *TokenInfo {
	if tokenStr == nil {
		return nil
	}

	claims, _ := m.ValidateToken(tokenStr)

	return &claims.TokenInfo
}
