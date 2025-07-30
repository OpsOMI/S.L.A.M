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

type TokenManager interface {
	GenerateToken(
		clientID, userID, username, nickname string,
		duration time.Duration,
	) (string, error)

	ValidateToken(
		tokenStr string,
	) (*Claims, error)

	ParseToken(
		tokenStr string,
	) (*TokenInfo, error)
}

type JWTManager struct {
	secret []byte
	issuer string
}

func NewJWTManager(secret string, issuer string) *JWTManager {
	return &JWTManager{
		secret: []byte(secret),
		issuer: issuer,
	}
}

func (jm *JWTManager) GenerateToken(clientID, userID, username, nickname string, duration time.Duration) (string, error) {
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
			Issuer:    jm.issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jm.secret)
}

func (jm *JWTManager) ValidateToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("jwt.unexpected_signing_method")
		}
		return jm.secret, nil
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

func (jm *JWTManager) ParseToken(tokenStr string) (*TokenInfo, error) {
	claims, err := jm.ValidateToken(tokenStr)
	if err != nil {
		return nil, err
	}

	return &claims.TokenInfo, nil
}
