package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {

	//Initdb()

	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())
	// Method:   GET
	// Resource: http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome Testing</h1>")
	})

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})
	app.Post("/news", func(ctx iris.Context) {
		news := News{}
		err := ctx.ReadJSON(&news)
		fmt.Println("news", news)
		if err != nil {
			panic(err.Error())
		}
		currentTime := time.Now()
		news.Created = currentTime.Format("2006-01-02 15:04:05")
		news.InsertNews()
		ctx.JSON(iris.Map{"message": "success"})
	})

	app.Run(iris.Addr(":3002"), iris.WithoutServerError(iris.ErrServerClosed))

}
