package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hthai2201/dw-go-23/exercises/06/common"
)

type PermissionStore interface {
	GetPermission() ([]interface{}, error)
}

func CheckPermission(sc common.AppContext, resourceName string, store PermissionStore) gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}
