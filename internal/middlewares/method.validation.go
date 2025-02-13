package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func MethodValidationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		allowedMethod := map[string]bool{
			http.MethodGet:   true,
			http.MethodPost:  true,
			http.MethodPut:   true,
			http.MethodPatch: true,
			http.MethodDelete: true,
			http.MethodOptions: true,
		}

		if !allowedMethod[c.Request().Method] {
			return c.JSON(http.StatusMethodNotAllowed, map[string]string{
				"error": "Method not allowed",
			})
		}

		return next(c)
	}
}