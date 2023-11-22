package storage

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func ConnectRabbitMQ() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Could not connect to RabbitMQ: %s", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Could not open a channel: %s", err)
	}
	return ch
}

func CreateTopicExchangeAndDeclareQueue() {
	ch := ConnectRabbitMQ()
	defer ch.Close()

	err := ch.ExchangeDeclare(
		"logs",  // name
		"topic", // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatalf("Could not declare exchange: %s", err)
	}

	q, err := ch.QueueDeclare(
		"",    // name (empty means random name)
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("Could not declare queue: %s", err)
	}

	err = ch.QueueBind(
		q.Name,   // queue name
		"logs.*", // binding key
		"logs",   // exchange name
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Could not bind queue to exchange: %s", err)
	}
	fmt.Println("Topic başarıyla oluşturuldu !")
}

func SendLogMessage(message string) {
	ch := ConnectRabbitMQ()
	defer ch.Close()
	// Bir mesaj oluştur
	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	}

	// Mesajı "logs" exchange'ine "logs.error" routing key'i ile publish et
	err := ch.Publish(
		"logs",       // exchange name
		"logs.error", // routing key
		false,        // mandatory
		false,
		msg,
	)
	if err != nil {
		log.Fatalf("Could not publish message: %s", err)
	}
}
func PublishEmailVerification(data []byte) {
	ch := ConnectRabbitMQ()
	defer ch.Close()
	// Bir mesaj oluştur
	msg := amqp.Publishing{
		ContentType: "application/json",
		Body:        data,
	}

	err := ch.Publish(
		"verifications",       // exchange name
		"verifications.email", // routing key
		false,                 // mandatory
		false,
		msg,
	)
	if err != nil {
		log.Fatalf("Could not publish message: %s", err)
	}
}
func CreateTopicExchangeVerificationAndDeclareQueue(ch *amqp.Channel) {

	// "logs" adında bir topic exchange oluştur
	err := ch.ExchangeDeclare(
		"verifications", // name
		"topic",         // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		log.Fatalf("Could not declare exchange: %s", err)
	}

	// Bir queue oluştur
	q, err := ch.QueueDeclare(
		"",    // name (empty means random name)
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("Could not declare queue: %s", err)
	}

	// Exchange ile queue arasında binding oluştur
	err = ch.QueueBind(
		q.Name,                // queue name
		"verifications.email", // binding key
		"verifications",       // exchange name
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Could not bind queue to exchange: %s", err)
	}
	fmt.Println("Verification Topic başarıyla oluşturuldu !")
}

func SendEmailVerificationToExchange(ch *amqp.Channel, message string) {
	// Bir mesaj oluştur
	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	}

	// Mesajı "logs" exchange'ine "logs.error" routing key'i ile publish et
	err := ch.Publish(
		"verifications",       // exchange name
		"verifications.email", // routing key
		false,                 // mandatory
		false,
		msg,
	)
	if err != nil {
		log.Fatalf("Could not publish message: %s", err)
	}
}
func SendSmsVerification(ch *amqp.Channel, message string) {
	// Bir mesaj oluştur
	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	}

	// Mesajı "logs" exchange'ine "logs.error" routing key'i ile publish et
	err := ch.Publish(
		"verifications",     // exchange name
		"verifications.sms", // routing key
		false,               // mandatory
		false,
		msg,
	)
	if err != nil {
		log.Fatalf("Could not publish message: %s", err)
	}
}
func CreateTopicExchangenotificationsAndDeclareQueue(ch *amqp.Channel) {

	// "logs" adında bir topic exchange oluştur
	err := ch.ExchangeDeclare(
		"notifications", // name
		"topic",         // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		log.Fatalf("Could not declare exchange: %s", err)
	}

	// Bir queue oluştur
	q, err := ch.QueueDeclare(
		"",    // name (empty means random name)
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("Could not declare queue: %s", err)
	}

	// Exchange ile queue arasında binding oluştur
	err = ch.QueueBind(
		q.Name,            // queue name
		"notifications.*", // binding key
		"notifications",   // exchange name
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Could not bind queue to exchange: %s", err)
	}
	fmt.Println("Verification Topic başarıyla oluşturuldu !")
}
func SendNotification(ch *amqp.Channel, routingKey, message string) {
	// Bir mesaj oluştur
	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	}

	// Mesajı "logs" exchange'ine "logs.error" routing key'i ile publish et
	err := ch.Publish(
		"notifications",             // exchange name
		"notifications."+routingKey, // routing key
		false,                       // mandatory
		false,
		msg,
	)
	if err != nil {
		log.Fatalf("Could not publish message: %s", err)
	}
}
