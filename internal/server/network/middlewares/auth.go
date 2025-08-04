package middlewares

import (
	"encoding/json"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/response"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/store"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/types"
)

func JWTAuthMiddleware(
	ts store.IJwtManager,
	targetRole ...string,
) types.Middleware {
	return func(next types.HandlerFunc) types.HandlerFunc {
		return func(conn net.Conn, args json.RawMessage, jwtToken *string) error {
			claims, err := ts.ValidateToken(jwtToken)
			if err != nil {
				return err
			}

			if len(targetRole) == 1 {
				if claims.Role != targetRole[0] {
					return response.Response(
						commons.StatusUnauthorized,
						"unauthorized: role mismatch",
						nil,
					)
				}
			}

			return next(conn, args, jwtToken)
		}
	}
}
