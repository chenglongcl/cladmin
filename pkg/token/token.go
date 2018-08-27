package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	// ErrMissingHeader means the `Authorization` header was empty.
	ErrMissingHeader  = errors.New("The length of the `Authorization` header is zero.")
	ErrRefreshExpired = errors.New("Token has refresh expired")
)

// Context is the context of the JSON web token.
type Context struct {
	ID       uint64
	Username string
}

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

// Parse validates the token with the specified secret,
// and returns the context if the token was valid.
func Parse(tokenString string, secret string, c *gin.Context) (*Context, error) {
	ctx := &Context{}

	// Parse the token.
	token, err := jwt.Parse(tokenString, secretFunc(secret))

	// Parse error.
	if err != nil {
		return ctx, err

		// Read the token if it's valid.
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		refreshExp := int64(claims["rexp"].(float64))
		if err := time.Now().Unix() < refreshExp; err != true {
			return ctx, ErrRefreshExpired
		}
		ctx.ID = uint64(claims["id"].(float64))
		ctx.Username = claims["username"].(string)
		c.Set("JWT_PAYLOAD", claims)
		c.Set("userId", ctx.ID)
		return ctx, nil

		// Other errors.
	} else {
		return ctx, err
	}
}

// ParseRequest gets the token from the header and
// pass it to the Parse function to parses the token.
func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")

	// Load the jwt secret from config
	secret := viper.GetString("jwt_secret")

	if len(header) == 0 {
		return &Context{}, ErrMissingHeader
	}

	var t string
	// Parse the header to get the token part.
	fmt.Sscanf(header, "Bearer %s", &t)
	return Parse(t, secret, c)
}

// Sign signs the context with the specified secret.
func Sign(ctx *gin.Context, c Context, secret string) (tokenString string, expiredAt string,
	refreshExpiredAt string, err error) {
	// Load the jwt secret from the Gin config if the secret isn't specified.
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}
	expConfig := viper.GetString("jwt_exp")
	m, _ := time.ParseDuration(expConfig + "s")
	RefreshExp := viper.GetString("jwt_refresh_exp")
	mt, _ := time.ParseDuration(RefreshExp + "s")
	now := time.Now()
	tokenExp := now.Add(m).Unix()
	tokenRefreshExp := now.Add(mt).Unix()
	// The token content.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       c.ID,
		"username": c.Username,
		"iat":      now.Unix(),
		"nbf":      now.Unix(),
		"sub":      c.ID,
		"exp":      tokenExp,
		"rexp":     tokenRefreshExp,
	})
	// Sign the token with the specified secret.
	tokenString, err = token.SignedString([]byte(secret))
	expiredAt = time.Unix(tokenExp, 0).Format("2006-01-02 15:04:05")
	refreshExpiredAt = time.Unix(tokenRefreshExp, 0).Format("2006-01-02 15:04:05")
	return
}

func ExtractClaims(c *gin.Context) jwt.MapClaims {
	claims, exists := c.Get("JWT_PAYLOAD")
	if !exists {
		return make(jwt.MapClaims)
	}

	return claims.(jwt.MapClaims)
}
