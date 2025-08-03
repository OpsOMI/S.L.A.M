package tokenstore

import (
	"time"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/response"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
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

type ITokenStore interface {
	GenerateToken(
		clientID, userID uuid.UUID,
		username, nickname, role string,
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

func (m *manager) GenerateToken(
	clientID, userID uuid.UUID,
	username, nickname, role string,
	duration time.Duration,
) (string, error) {
	claims := Claims{
		TokenInfo: TokenInfo{
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

func (m *manager) ValidateToken(tokenStr *string) (*Claims, error) {
	if tokenStr == nil {
		return nil, response.Response(commons.StatusUnauthorized, "JWT token is required", nil)
	}

	token, err := jwt.ParseWithClaims(*tokenStr, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, response.Response(commons.StatusUnauthorized, "Unexpected signing method", nil)
		}
		return m.secret, nil
	})
	if err != nil {
		return nil, response.Response(commons.StatusUnauthorized, "Invalid JWT token", nil)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, response.Response(commons.StatusUnauthorized, "Invalid JWT token", nil)
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
