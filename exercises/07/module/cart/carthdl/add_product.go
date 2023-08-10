package carthdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hthai2201/dw-go-23/exercises/07/common"
	"github.com/hthai2201/dw-go-23/exercises/07/module/cart/cartmodel"
	"github.com/hthai2201/dw-go-23/exercises/07/module/cart/cartrepo"
	"github.com/hthai2201/dw-go-23/exercises/07/module/cart/cartstorage"
)

func AddCartProduct(appCtx common.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var data cartmodel.CartProductAdd

		if err := c.ShouldBind(&data); err != nil && err.Error() != "EOF" {
			panic(err)
		}

		db := appCtx.GetDBConnection()
		store := cartstorage.NewMysql(db)
		repo := cartrepo.NewAddCartProductRepo(store)
		if err := repo.AddCartProduct(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
