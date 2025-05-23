package tokenizer

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	errGenerateToken = errors.New("token generation failed")
	errParseToken    = errors.New("token parsing failed")
	errVerifyToken   = errors.New("token verification failed")
)

type Tokenizer interface {
	GenerateToken(role string) (*string, error)
	VerifyToken(tokenString string) (*jwt.Token, error)
}

type tokenizer struct {
	tokenIssuer string
	secretKey   []byte
}

func New(iss, key string) Tokenizer {
	return &tokenizer{
		tokenIssuer: iss,
		secretKey:   []byte(key),
	}
}

func (t *tokenizer) GenerateToken(role string) (*string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": role,
		"iss": t.tokenIssuer,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	})

	token, err := claims.SignedString(t.secretKey)
	if err != nil {
		return nil, errGenerateToken
	}

	return &token, nil
}

func (t *tokenizer) VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return t.secretKey, nil
	})
	if err != nil {
		return nil, errParseToken
	}

	if !token.Valid {
		return nil, errVerifyToken
	}

	return token, nil
}
