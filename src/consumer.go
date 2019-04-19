package main

import (
	"fmt"
	"time"
	"log"
	"github.com/streadway/amqp"
)

//var rabbitmq_server = "rabbit"
var rabbitmq_server = "localhost"
var rabbitmq_port = "5672"
var rabbitmq_queue = "sensordata"
var rabbitmq_user = "guest"
var rabbitmq_pass = "guest"

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}


// Receive Order from Queue to Process
func main() {

	fmt.Println("Consumer Starts")
	fmt.Println("current time: ", time.Now())
	conn, err := amqp.Dial("amqp://"+rabbitmq_user+":"+rabbitmq_pass+"@"+rabbitmq_server+":"+rabbitmq_port+"/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		rabbitmq_queue, // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	sensorread := make(chan string)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	<-sensorread

fmt.Println("Consumer Ends")


	
}


