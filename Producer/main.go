package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/LoJiPhin/kumparan/Contracts"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/streadway/amqp"
)

var (
	amqpURI = flag.String("amqp", "amqp://rabbitmq:rabbitmq@rabbit1:5672/", "AMQP URI")
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
func init() {
	flag.Parse()
	initAmqp()
}

var conn *amqp.Connection
var ch *amqp.Channel

func publishMessages(news Contracts.News) {
	payload, err := json.Marshal(news)
	failOnError(err, "Failed to marshal JSON")

	err = ch.Publish(
		"go-test-exchange", // exchange
		"go-test-key",      // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Transient,
			ContentType:  "application/json",
			Body:         payload,
			Timestamp:    time.Now(),
		})

	failOnError(err, "Failed to Publish on RabbitMQ")
}

func initAmqp() {
	var err error

	conn, err = amqp.Dial(*amqpURI)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")

	err = ch.ExchangeDeclare(
		"test-exchange", // name
		"direct",        // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // noWait
		nil,             // arguments
	)
	failOnError(err, "Failed to declare the Exchange")
}

func main() {

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
		// news := News{}
		// err := ctx.ReadJSON(&news)
		// fmt.Println("news", news)
		// if err != nil {
		// 	panic(err.Error())
		// }
		// currentTime := time.Now()
		// news.Created = currentTime.Format("2006-01-02 15:04:05")
		// news.InsertNews()
		// ctx.JSON(iris.Map{"message": "success"})
		news := Contracts.News{}
		err := ctx.ReadJSON(&news)
		failOnError(err, "Failed to Read Json")
		publishMessages(news)

	})

	app.Run(iris.Addr(":3002"), iris.WithoutServerError(iris.ErrServerClosed))

}
