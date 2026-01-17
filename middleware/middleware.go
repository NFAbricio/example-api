package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/NFAbricio/example-api/users"
)

type UserMiddleware struct{}

func NewUserMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}

func (um *UserMiddleware) Middleware() gin.HandlerFunc{
	return func (httpContext *gin.Context)  {
		secret := []byte(viper.GetString("JWT_SECRET"))

		cookie, err := httpContext.Cookie("token")
		if err != nil {
			httpContext.JSON(http.StatusUnauthorized, gin.H{"error": "cookie not found"})
			httpContext.Abort()
			return
		}

		if cookie == "" {
			httpContext.JSON(http.StatusUnauthorized, gin.H{"error": "token not found"})
			httpContext.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(
			cookie,
			&users.ClaimsUser{},
			func(token *jwt.Token) (interface{}, error) {
				return secret, nil
			},
		)
		if err != nil {
			httpContext.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			httpContext.Abort()
			return
		}

		claims, ok := token.Claims.(*users.ClaimsUser)
		if !ok || !token.Valid {
			httpContext.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			httpContext.Abort()
			return
		}

		if claims.Role != (viper.GetString("USER_ROLE_SECRET")) {
			httpContext.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			httpContext.Abort()
			return
		}

		httpContext.Set((viper.GetString("USER_ROLE_SECRET")), claims.User)
		httpContext.Set("isAuthenticated", true)

		httpContext.Next()
	}
}

func (um *UserMiddleware) GetUserFromMiddleware(httpContext *gin.Context) (*users.ClaimsUser, error) {
	secret := []byte(viper.GetString("JWT_SECRET"))

	cookie, err := httpContext.Cookie("token")
	if err != nil {
		httpContext.JSON(http.StatusUnauthorized, gin.H{"error": "cookie not found"})
		httpContext.Abort()
		return nil, fmt.Errorf("cookie not found: %w", err)
	}

	token, err := jwt.ParseWithClaims(
		cookie,
		&users.ClaimsUser{},
		func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		},
	)

	if err != nil {
		httpContext.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		httpContext.Abort()
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*users.ClaimsUser)
	if !ok || !token.Valid {
		httpContext.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		httpContext.Abort()
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	return claims, nil
}