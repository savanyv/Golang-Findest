package middlewares

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/savanyv/Golang-Findest/internal/helpers"
)

func AuthMiddlewares(jwtService helpers.JWTService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(401, "Unauthorized")
			}

			bearerToken := strings.Split(authHeader, " ")
			if len(bearerToken) != 2 {
				return echo.NewHTTPError(401, "Unauthorized")
			}

			token := bearerToken[1]
			claims, err := jwtService.ValidateToken(token)
			if err != nil {
				return echo.NewHTTPError(401, "Unauthorized")
			}

			c.Set("userID", claims.UserID)
			c.Set("email", claims.Email)

			return next(c)
		}
	}
}