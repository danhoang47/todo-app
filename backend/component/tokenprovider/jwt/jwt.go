package jwt

import (
	"time"
	"todo/component/tokenprovider"

	"github.com/dgrijalva/jwt-go"
)

type jwtProvider struct {
	secrect string
}

func NewJwtProvider(secret string) *jwtProvider {
	return &jwtProvider{secret}
}

type myClaims struct {
	Payload tokenprovider.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
	})

	token, err := t.SignedString([]byte(j.secrect))

	if err != nil {
		return nil, tokenprovider.ErrEncodingToken
	}

	return &tokenprovider.Token{Token: token, Created: time.Now(), Expiry: expiry}, nil
}

func (j *jwtProvider) Validate(token string) (*tokenprovider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(token, &myClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(j.secrect), nil
	})

	if err != nil {
		return nil, tokenprovider.ErrInvalidToken
	}

	if !res.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}

	claims, ok := res.Claims.(*myClaims)
	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}

	return &claims.Payload, nil
}
