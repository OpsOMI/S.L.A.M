package store

import (
	"time"

	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/response"
	"github.com/OpsOMI/S.L.A.M/internal/shared/store"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type IJwtManager interface {
	GenerateToken(
		clientID, userID uuid.UUID,
		username, nickname, role string,
		duration time.Duration,
	) (string, error)

	ValidateToken(
		tokenStr *string,
	) (*store.Claims, error)

	ParseToken(
		tokenStr *string,
	) *store.TokenInfo
}

type manager struct {
	issuer string
	secret []byte
}

func NewManager(
	issuer, secret string,
) IJwtManager {
	return &manager{
		issuer: issuer,
		secret: []byte(secret),
	}
}

func (m *manager) GenerateToken(
	clientID, userID uuid.UUID,
	username, nickname, role string,
	duration time.Duration,
) (string, error) {
	claims := store.Claims{
		TokenInfo: store.TokenInfo{
			ClientID: clientID,
			UserID:   userID,
			Username: username,
			Nickname: nickname,
			Role:     role,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    m.issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(m.secret)
	if err != nil {
		return "", response.Response(
			commons.StatusInternalServerError,
			"Failed to sign JWT token",
			nil,
		)
	}

	return signedToken, nil
}

func (m *manager) ValidateToken(tokenStr *string) (*store.Claims, error) {
	if tokenStr == nil {
		return nil, response.Response(commons.StatusUnauthorized, "JWT token is required", nil)
	}
	if *tokenStr == "" {
		return nil, response.Response(commons.StatusUnauthorized, "Unauthorized", nil)
	}

	token, err := jwt.ParseWithClaims(*tokenStr, &store.Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, response.Response(commons.StatusUnauthorized, "Unexpected signing method", nil)
		}
		return m.secret, nil
	})
	if err != nil {
		return nil, response.Response(commons.StatusUnauthorized, "Invalid JWT token", nil)
	}

	claims, ok := token.Claims.(*store.Claims)
	if !ok || !token.Valid {
		return nil, response.Response(commons.StatusUnauthorized, "Invalid JWT token", nil)
	}
	return claims, nil
}

// This function assumes the JWT is already validated by middleware, so it skips error handling and returns the parsed token info directly.
func (m *manager) ParseToken(tokenStr *string) *store.TokenInfo {
	if tokenStr == nil {
		return nil
	}

	claims, _ := m.ValidateToken(tokenStr)

	return &claims.TokenInfo
}
