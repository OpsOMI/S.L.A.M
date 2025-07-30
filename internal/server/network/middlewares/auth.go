package middlewares

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/response"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/tokenstore"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/types"
)

func JWTAuthMiddleware(ts tokenstore.ITokenStore) types.Middleware {
	return func(next types.HandlerFunc) types.HandlerFunc {
		return func(conn net.Conn, args json.RawMessage, jwtToken *string) error {
			_, err := ts.ValidateToken(jwtToken)
			if err != nil {
				response.Forbidden(conn)
				return fmt.Errorf("forbidden: %s", err)
			}

			return next(conn, args, jwtToken)
		}
	}
}
