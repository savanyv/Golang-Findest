package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/savanyv/Golang-Findest/internal/config"
)

type JWTClaim struct {
	UserID string `json:"user_id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type JWTService interface {
	GenerateToken(userID uint, email string) (string, error)
	ValidateToken(tokenString string) (*JWTClaim, error)
}

type jwtService struct {
	secretKey string
	issuer string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: config.LoadConfig().Jwt.SecretKey,
		issuer: "findest",
	}
}

func (j *jwtService) GenerateToken(userID uint, email string) (string, error) {
	claims := &JWTClaim{
		UserID: string(rune(userID)),
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: j.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtService) ValidateToken(tokenString string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok || !token.Valid {
		return nil, jwt.ErrInvalidKey
	}

	return claims, nil
}

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.NewHTTPError(401, "Missing token")
		}

		jwtService := NewJWTService()
		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			return echo.NewHTTPError(401, "Invalid token")
		}

		c.Set("user", claims)
		return next(c)
	}
}