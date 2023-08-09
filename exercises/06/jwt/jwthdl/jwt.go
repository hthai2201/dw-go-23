package jwthdl

import (
	"context"

	"github.com/hthai2201/dw-go-23/exercises/06/common"
	"github.com/hthai2201/dw-go-23/exercises/06/module/user/usermodel"
	"github.com/hthai2201/dw-go-23/exercises/06/token"
)

type JwtRepo interface {
	Validate(ctx context.Context, payload *token.JwtPayload) (*common.SimpleUser, error)
}

type jwtHdl struct {
	repo JwtRepo
}

func NewJwtHdl(repo JwtRepo) *jwtHdl {
	return &jwtHdl{
		repo: repo,
	}
}

func (hdl *jwtHdl) Validate(ctx context.Context, payload *token.JwtPayload) (*common.SimpleUser, error) {
	user, err := hdl.repo.Validate(ctx, payload)
	if err != nil {
		return nil, common.ErrCannotGetEntity(usermodel.EntityName, err)
	}
	return user, nil
}
