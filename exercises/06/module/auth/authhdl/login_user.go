package authhdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hthai2201/dw-go-23/exercises/06/common"
	"github.com/hthai2201/dw-go-23/exercises/06/module/auth/authmodel"
	"github.com/hthai2201/dw-go-23/exercises/06/module/auth/authrepo"
	"github.com/hthai2201/dw-go-23/exercises/06/module/user/userstorage"
	"github.com/hthai2201/dw-go-23/exercises/06/token"
	"github.com/hthai2201/dw-go-23/exercises/06/token/jwt"
)

func Login(appCtx common.AppContext, secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {

		var loginUserData authmodel.LoginUser

		if err := c.ShouldBind(&loginUserData); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		db := appCtx.GetDBConnection()
		tokenProvider := jwt.NewTokenProvider(token.WithSecretKey([]byte(secretKey)))

		store := userstorage.NewUserMysql(db)
		repo := authrepo.NewLoginUserRepo(store, tokenProvider)
		account, err := repo.LoginUser(c.Request.Context(), &loginUserData)

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"access_token": account.AccessToken,
		})
	}
}
