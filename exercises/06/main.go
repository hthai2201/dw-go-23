package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/hthai2201/dw-go-23/exercises/06/appctx"
	"github.com/hthai2201/dw-go-23/exercises/06/middleware"
	"github.com/hthai2201/dw-go-23/exercises/06/module/auth/authhdl"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Tier layer

// Repo (business logic) -----> Storage

func main() {
	dbConStr := os.Getenv("DBConnStr")
	secretKey := os.Getenv("SECRET_KEY")
	db, err := gorm.Open(mysql.Open(dbConStr), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

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
	r.Run(":" + "8000")
}

type Requester interface {
	UserId() int
	Role() string
	FirstName() string
	LastName() string
}

func checkClosure() {
	arr := make([]func(), 10)

	for i := 0; i <= 9; i++ {
		f := func(y int) func() {
			// y is value of i
			return func() {
				log.Println(y) // pointer to y, because closure capture all variable outside as a pointer
			}
		}

		arr[i] = f(i + 2)
	}

	for i := range arr {
		arr[i]()
	}
}
