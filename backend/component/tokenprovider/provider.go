package tokenprovider

import (
	"errors"
	"time"
	"todo/common"
)

type Provider interface {
	Generate(data TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

var (
	ErrNotFound = common.NewCustomerError(
		errors.New("token not found"),
		"token not found",
		"ErrNotFound",
	)

	ErrEncodingToken = common.NewCustomerError(
		errors.New("error encoding the token"),
		"error encoding the token",
		"ErrEncodingToken",
	)

	ErrInvalidToken = common.NewCustomerError(
		errors.New("invalid token provided"),
		"invalid token provided",
		"ErrInvalidToken",
	)
)

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created_at"`
	Expiry  int       `json:"expiry"`
}

type TokenPayload struct {
	UserId uint   `json:"userId"`
	Role   string `json:"role"`
}
