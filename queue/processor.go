package queue

import (
	"github.com/spf13/viper"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type config struct {
	ConnectionString string `yaml:"rabbitmq.connectionString"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func connect() (*amqp.Channel, amqp.Queue) {
	settings := viper.New()
	settings.SetConfigName("settings")
	settings.AddConfigPath(".")
	err := settings.ReadInConfig()
	failOnError(err, "Unable to read config")

	conf := &config{}
	err = settings.Unmarshal(conf)
	failOnError(err, "Unable to read config")

	fmt.Printf("Connection String: %s", conf.ConnectionString)

	conn, err := amqp.Dial(settings.GetString("rabbitmq.connectionString"))
	failOnError(err, "Unable to connect")

	ch, err := conn.Channel()
	failOnError(err, "Unable to create channel")

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Unable to declare queue")

	return ch, q
}

func SendQueue() {
	ch, q := connect()

	body := "FUCKING SHIT!!"

	err := ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Unable to send message")
}

func ReadQueue() {
	ch, q := connect()

	msg, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Unable to consume")

	forever := make(chan bool)

	go func() {
		for d := range msg {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
