package token

import (
	"errors"
	"time"

	"github.com/hthai2201/dw-go-23/exercises/06/common"
	"github.com/hthai2201/dw-go-23/exercises/06/module/user/usermodel"
)

var (
	ErrNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"ErrNotFound",
	)
	ErrEncodingToken = common.NewCustomError(errors.New("error encoding the token"),
		"error encoding the token",
		"ErrEncodingToken",
	)
	ErrInvalidToken = common.NewCustomError(errors.New("invalid token provided"),
		"invalid token provided",
		"ErrInvalidToken",
	)
)

type JwtPayload struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
}

func (jp *JwtPayload) GetUserId() int {
	return jp.UserId
}

type Provider interface {
	Generate(user usermodel.User, opts ...GenerateOption) (*Token, error)
	Inspect(token string) (*JwtPayload, error)
	String() string
}

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  time.Time `json:"expiry"`
}
