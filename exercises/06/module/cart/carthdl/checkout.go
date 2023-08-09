package carthdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hthai2201/dw-go-23/exercises/06/common"
	"github.com/hthai2201/dw-go-23/exercises/06/module/cart/cartrepo"
	"github.com/hthai2201/dw-go-23/exercises/06/module/cart/cartstorage"
)

func CheckoutCart(appCtx common.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {

		db := appCtx.GetDBConnection()
		store := cartstorage.NewMysql(db)
		repo := cartrepo.NewCheckoutCartRepo(store)
		checkout, err := repo.Checkout(c.Request.Context())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, checkout)
	}
}
