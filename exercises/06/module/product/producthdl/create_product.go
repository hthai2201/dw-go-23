package producthdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hthai2201/dw-go-23/exercises/06/common"
	"github.com/hthai2201/dw-go-23/exercises/06/module/product/productmodel"
	"github.com/hthai2201/dw-go-23/exercises/06/module/product/productrepo"
	"github.com/hthai2201/dw-go-23/exercises/06/module/product/productstorage"
)

func CreateProduct(appCtx common.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var data productmodel.ProductCreate

		if err := c.ShouldBind(&data); err != nil && err.Error() != "EOF" {
			panic(err)
		}

		db := appCtx.GetDBConnection()
		store := productstorage.NewMysql(db)
		repo := productrepo.NewCreateProductRepo(store)
		if err := repo.CreateProduct(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
