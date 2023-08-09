package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hthai2201/dw-go-23/exercises/06/common"
	"github.com/hthai2201/dw-go-23/exercises/06/jwt/jwthdl"
	"github.com/hthai2201/dw-go-23/exercises/06/jwt/jwtrepo"
	"github.com/hthai2201/dw-go-23/exercises/06/module/user/userstorage"
	"github.com/hthai2201/dw-go-23/exercises/06/token"
	"github.com/hthai2201/dw-go-23/exercises/06/token/jwt"
)

type jwtMiddleware struct {
	TokenType string
	jwthdl    jwthdl.JwtRepo
}

func NewJwtMiddleware(tokType string, hdl jwthdl.JwtRepo) *jwtMiddleware {
	return &jwtMiddleware{
		TokenType: tokType,
		jwthdl:    hdl,
	}
}

func (jm *jwtMiddleware) FromAuthHeaderAsBearerToken(c *gin.Context, key string) (tok string, err error) {
	authHeader := c.Request.Header.Get(key)
	if authHeader == "" {
		return "", errors.New("empty auth header")
	}

	parts := strings.SplitN(authHeader, " ", 2)

	if !(len(parts) == 2 && parts[0] == jm.TokenType) {
		return "", errors.New("invalid auth header")
	}
	return parts[1], nil
}

func RequiredAuth(appCtx common.AppContext, secretKey string) gin.HandlerFunc {
	tokProvider := jwt.NewTokenProvider(token.WithSecretKey([]byte(secretKey)))
	return func(c *gin.Context) {

		store := userstorage.NewUserMysql(appCtx.GetDBConnection())
		repo := jwtrepo.NewJwtVerifyRepo(store)
		hdl := jwthdl.NewJwtHdl(repo)

		jwtMidd := NewJwtMiddleware("Bearer", hdl)

		tok, err := jwtMidd.FromAuthHeaderAsBearerToken(c, "Authorization")
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		payload, err := tokProvider.Inspect(tok)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		sUser, err := jwtMidd.jwthdl.Validate(c.Request.Context(), payload)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		c.Set(common.KeyCurrentUser, sUser)

		c.Next()
	}
}
