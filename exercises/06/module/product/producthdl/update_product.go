package producthdl

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hthai2201/dw-go-23/exercises/06/common"
	"github.com/hthai2201/dw-go-23/exercises/06/module/product/productmodel"
	"github.com/hthai2201/dw-go-23/exercises/06/module/product/productrepo"
	"github.com/hthai2201/dw-go-23/exercises/06/module/product/productstorage"
)

func UpdateProduct(appCtx common.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("product-id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}
		var data productmodel.ProductUpdate
		if err := c.ShouldBind(&data); err != nil && err.Error() != "EOF" {
			panic(err)
		}
		data.ID = id

		db := appCtx.GetDBConnection()
		store := productstorage.NewMysql(db)
		repo := productrepo.NewUpdateProductRepo(store)
		if err := repo.UpdateProduct(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
