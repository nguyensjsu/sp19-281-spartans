package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var rabbitmq_server = "rabbit"
var rabbitmq_port = "5672"
var rabbitmq_queue = "ha.spartans"
var rabbitmq_user = "guest"
var rabbitmq_pass = "guest"
var producer = "producer1"

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func queue_send(message string) {
	conn, err := amqp.Dial("amqp://" + rabbitmq_user + ":" + rabbitmq_pass + "@" + rabbitmq_server + ":" + rabbitmq_port + "/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		rabbitmq_queue, // name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
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

func getCPUSample() (idle, total uint64) {
	contents, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range (lines) {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}
				total += val // tally up all the numbers to get total ticks
				if i == 4 { // idle is the 5th field in the cpu line
					idle = val
				}
			}
			return
		}
	}
	return
}
func getTemp() (temp float64) {
	contents, err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		return
	}
	tempStr := strings.Split(string(contents), "\n")[0]
	tempFloat, _ := strconv.ParseFloat(tempStr, 64)

	temp = tempFloat / float64(1000)
	return
}

func getMemory() (memPct float64) {
	out, _ := exec.Command("vmstat", "-s").Output()
	lines := strings.Split(string(out), "\n")
	totalMemory := strings.Split(strings.Trim(lines[0], " "), " ")[0]
	totalMemoryFloat, _ := strconv.ParseFloat(totalMemory, 64)
	usedMemory := strings.Split(strings.Trim(lines[1], " "), " ")[0]
	usedMemoryFloat, _ := strconv.ParseFloat(usedMemory, 64)
	memPct = usedMemoryFloat * float64(100) / totalMemoryFloat
	return
}

func main() {
	idle0, total0 := getCPUSample()

	idleTicks := float64(idle0)
	totalTicks := float64(total0)
	cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks

	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))
	message := fmt.Sprint(producer, ",", time.Now(), ",", cpuUsage, ",", getMemory(), ",", getTemp())
	queue_send(message)
}
