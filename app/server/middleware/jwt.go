package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/iqbalrestu07/datting-apps-api/domain"
	"github.com/labstack/echo/v4"
)

type AuthCustomClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// jwt service
type JWTServiceInterface interface {
	GenerateToken(domain.User) string
	ValidateToken(token string) (*jwt.Token, error)
	Claims(ctx echo.Context) AuthCustomClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

// auth-jwt
func NewJWTAuthService() JWTServiceInterface {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "datting-app",
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (s *jwtService) GenerateToken(user domain.User) string {
	claims := &AuthCustomClaims{
		UserID: user.ID.String(),
		StandardClaims: jwt.StandardClaims{
			Issuer:   s.issuer,
			IssuedAt: time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
}

func (s *jwtService) Claims(c echo.Context) (claims AuthCustomClaims) {
	const BEARER_SCHEMA = "Bearer"
	authHeader := c.Request().Header.Get("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA)+1:]

	token, _ := s.ValidateToken(tokenString)

	byt, _ := json.Marshal(token.Claims)
	err := json.Unmarshal(byt, &claims)
	if err != nil {
		return claims
	}

	return claims
}

func Authorize(s JWTServiceInterface) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			const BEARER_SCHEMA = "Bearer"
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Need Token to Get Resources")
			}

			tokenString := authHeader[len(BEARER_SCHEMA)+1:]
			token, err := s.ValidateToken(tokenString)

			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
			}

			if !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
			}

			claim := s.Claims(c)

			c.Set("user_id", claim.UserID)

			return next(c)
		}
	}
}
