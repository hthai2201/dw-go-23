package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hthai2201/dw-go-23/exercises/07/module/user/usermodel"
	"github.com/hthai2201/dw-go-23/exercises/07/token"
)

type authClaims struct {
	Payload token.JwtPayload `json:"payload"`
	jwt.StandardClaims
}

type JWT struct {
	opts token.Options
}

func NewTokenProvider(opts ...token.Option) token.Provider {
	return &JWT{
		opts: token.NewOptions(opts...),
	}
}

func (j *JWT) Generate(user usermodel.User, opts ...token.GenerateOption) (*token.Token, error) {
	// parse the options
	options := token.NewGenerateOptions(opts...)

	// generate the JWT
	expiry := time.Now().Add(options.Expiry)
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaims{
		token.JwtPayload{
			UserId: user.ID,
			//Role: account.User.Roles.Value(),
		},
		jwt.StandardClaims{
			ExpiresAt: expiry.Unix(),
		},
	})
	tok, err := t.SignedString(j.opts.SecretKey)
	if err != nil {
		return nil, err
	}

	// return the token
	return &token.Token{
		Token:   tok,
		Expiry:  expiry,
		Created: time.Now(),
	}, nil
}

func (j *JWT) Inspect(t string) (*token.JwtPayload, error) {

	// parse the public key
	res, err := jwt.ParseWithClaims(t, &authClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.opts.SecretKey, nil
	})

	if err != nil {
		return nil, token.ErrInvalidToken
	}

	// validate the token
	if !res.Valid {
		return nil, token.ErrInvalidToken
	}
	claims, ok := res.Claims.(*authClaims)
	if !ok {
		return nil, token.ErrInvalidToken
	}

	// return the token
	return &claims.Payload, nil
}

func (j *JWT) String() string {
	panic("implement me")
}
