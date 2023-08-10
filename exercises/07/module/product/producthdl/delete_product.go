package producthdl

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hthai2201/dw-go-23/exercises/07/common"
	"github.com/hthai2201/dw-go-23/exercises/07/module/product/productrepo"
	"github.com/hthai2201/dw-go-23/exercises/07/module/product/productstorage"
)

// DELETE v1/products/:product-id

func DeleteProduct(appCtx common.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("product-id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		db := appCtx.GetDBConnection()

		store := productstorage.NewMysql(db)
		repo := productrepo.NewDeleteProductRepo(store)

		_, err = repo.DeleteProduct(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
