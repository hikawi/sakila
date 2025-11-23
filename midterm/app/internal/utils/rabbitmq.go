package utils

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func InitRabbitMQ(queueName string) amqp.Queue {
	conn, err := amqp.Dial(FatalEnv("RABBITMQ_URL"))
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	FailOnError(err, "Failed to declare a queue")
	return q
}
