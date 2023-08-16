package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/hthai2201/dw-go-23/exercises/07/appctx"
	"github.com/hthai2201/dw-go-23/exercises/07/middleware"
	"github.com/hthai2201/dw-go-23/exercises/07/module/auth/authhdl"
	"github.com/hthai2201/dw-go-23/exercises/07/module/cart/carthdl"
	"github.com/hthai2201/dw-go-23/exercises/07/module/cart/cartmodel"
	"github.com/hthai2201/dw-go-23/exercises/07/module/product/producthdl"
	"github.com/hthai2201/dw-go-23/exercises/07/module/product/productmodel"
	"github.com/hthai2201/dw-go-23/exercises/07/module/user/usermodel"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Tier layer

// Repo (business logic) -----> Storage

func main() {
	dbConStr := os.Getenv("DBConnStr")
	secretKey := os.Getenv("SECRET_KEY")
	db, err := gorm.Open(postgres.Open(dbConStr), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&usermodel.User{}, &productmodel.Product{}, &cartmodel.Cart{})
	appCtx := appctx.NewAppContext(db.Debug())

	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	v1 := r.Group("/v1")

	auth := v1.Group("/auth")
	auth.POST("/register", authhdl.Register(appCtx))
	auth.POST("/login", authhdl.Login(appCtx, secretKey))

	productRoutes := v1.Group("/products")
	// productRoutes.Use(middleware.RequiredAuth(appCtx, secretKey))
	productRoutes.GET("", producthdl.GetListProducts(appCtx))
	productRoutes.POST("", producthdl.CreateProduct(appCtx))
	productRoutes.DELETE("/:product-id", producthdl.DeleteProduct(appCtx))
	productRoutes.PUT("/:product-id", producthdl.UpdateProduct(appCtx))

	cartRoutes := v1.Group("/cart")
	// cartRoutes.Use(middleware.RequiredAuth(appCtx, secretKey))
	cartRoutes.GET("", carthdl.GetListCartProducts(appCtx))
	cartRoutes.POST("/add", carthdl.AddCartProduct(appCtx))
	cartRoutes.DELETE("/remove", carthdl.RemoveCartProduct(appCtx))
	cartRoutes.POST("/checkout", carthdl.CheckoutCart(appCtx))

	r.Run(":" + "8080")
}
