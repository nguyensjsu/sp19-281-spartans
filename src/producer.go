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


// Send Order to Queue for Processing
func queue_send(message string) {
	conn, err := amqp.Dial("amqp://"+rabbitmq_user+":"+rabbitmq_pass+"@"+rabbitmq_server+":"+rabbitmq_port+"/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		rabbitmq_queue, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := message
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}

func main(){
	
	fmt.Println("Producer Starts")
	fmt.Println("current time: ", time.Now())


	for i:=0; i < 4 ; i++{

        fmt.Println("Sending data to queue")
		queue_send("testmessage")
	}

	fmt.Println("Producer Ends")
}