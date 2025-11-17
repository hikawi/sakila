package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

type WithAMQP struct {
	Channel *amqp.Channel
	Queue   amqp.Queue
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func (rabbit *WithAMQP) SendSakura(g *gin.Context) {
	var body struct {
		Message string `json:"message"`
	}
	err := g.ShouldBindBodyWithJSON(&body)
	if err != nil {
		g.AbortWithStatusJSON(400, gin.H{"message": "bad request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = rabbit.Channel.PublishWithContext(ctx, "", "bff", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body.Message),
	})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body.Message)
}

func main() {
	server := gin.Default()

	conn, err := amqp.Dial("amqp://guest:guest@awad-week7-rabbitmq:5672")
	failOnError(err, "Can't establish rabbitmq connection")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"bff", // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	withAMQP := WithAMQP{Channel: ch, Queue: q}

	server.POST("/", withAMQP.SendSakura)

	server.Run(":80")
}
