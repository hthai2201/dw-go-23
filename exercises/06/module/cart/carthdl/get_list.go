package carthdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hthai2201/dw-go-23/exercises/06/common"
	"github.com/hthai2201/dw-go-23/exercises/06/module/cart/cartmodel"
	"github.com/hthai2201/dw-go-23/exercises/06/module/cart/cartrepo"
	"github.com/hthai2201/dw-go-23/exercises/06/module/cart/cartstorage"
)

func GetListCartProducts(appCtx common.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var p cartmodel.ListParam

		if err := c.ShouldBind(&p); err != nil && err.Error() != "EOF" {
			panic(err)
		}

		p.Fulfill()

		db := appCtx.GetDBConnection()

		store := cartstorage.NewMysql(db)
		repo := cartrepo.NewListStorage(store)

		result, err := repo.List(c.Request.Context(), &p.Paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, p.Paging, p.ListFilter))
	}
}
