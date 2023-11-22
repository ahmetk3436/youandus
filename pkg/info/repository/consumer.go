package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"
	"youandus/internal/storage"

	"github.com/go-mail/mail"
	"github.com/streadway/amqp"
)

func ConsumeVerification() {
	ch := storage.ConnectRabbitMQ()
	defer ch.Close()
	storage.CreateTopicExchangeVerificationAndDeclareQueue(ch)

	// Bir queue oluştur
	q, err := ch.QueueDeclare(
		"verification_queue", // name (empty means random name)
		false,                // durable
		false,                // delete when unused
		true,                 // exclusive
		false,                // no-wait
		nil,                  // arguments
	)
	if err != nil {
		log.Fatalf("Could not declare queue: %s", err)
	}

	err = ch.QueueBind(
		q.Name,                // queue name
		"verifications.email", // routing key
		"verifications",       // exchange name
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Could not bind queue to exchange: %s", err)
	}

	// Mesajları işleme fonksiyonunu tanımla
	processMessage := func(msg amqp.Delivery) {
		// Mesajı doğrulama işlemine tabi tut
		var user User
		err := json.Unmarshal(msg.Body, &user)
		if err != nil {
			log.Fatalf("Could not decode JSON: %s", err)
		}

		// Doğrulama işlemini gerçekleştir
		SendEmail(user.Email)

		fmt.Printf("Verification code received: %d %s\n", user.ID, user.Email)

		// Mesajı acknowledge et
		msg.Ack(false)
	}

	// Queue'dan mesajları tüket
	msgs, err := ch.Consume(
		q.Name, // queue name
		"",     // consumer name (empty means random name)
		false,  // auto-acknowledge (disable auto-ack)
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		log.Fatalf("Could not consume messages: %s", err)
	}

	// Mesajları işleme döngüsü
	for msg := range msgs {
		processMessage(msg)
	}
}

type User struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

func ConsumeNotification() {
	ch := storage.ConnectRabbitMQ()
	defer ch.Close()
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

	// Queue'dan mesajları tüket
	msgs, err := ch.Consume(
		q.Name, // queue name
		"",     // consumer name (empty means random name)
		true,   // auto-acknowledge
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		log.Fatalf("Could not consume messages: %s", err)
	}

	// Mesajları işle
	go func() {
		for msg := range msgs {
			// Mesajı bildirim işlemine tabi tut
			routingKey := msg.RoutingKey
			notification := string(msg.Body)
			// Bildirim işlemini gerçekleştir
			// ...
			fmt.Printf("Notification received with routing key '%s': %s\n", routingKey, notification)
		}
	}()
}

func SendEmail(userEmail string) error {
	verificationCode := GetEmailVerificationCode(userEmail)

	// Belirli bir deneme sayısı ile verificationCode alınır
	maxAttempts := 5
	attempts := 0
	for verificationCode == "" && attempts < maxAttempts {
		time.Sleep(5 * time.Second) // 5 saniye ara ver
		verificationCode = GetEmailVerificationCode(userEmail)
		attempts++
	}

	if verificationCode == "" {
		return errors.New("verification code could not be obtained")
	}

	m := mail.NewMessage()
	m.SetHeader("From", "testerahmet16@gmail.com")
	m.SetHeader("To", "testerahmet16@gmail.com")
	m.SetHeader("Subject", "Verification Code")
	m.SetBody("text/html", "Please click the link and verify your account !"+"\n"+" Link : "+"https://api.youandus.net/api/verifyemail?email="+userEmail+"&verification_code="+verificationCode)

	d := mail.NewDialer("smtp.gmail.com", 587, "testerahmet16@gmail.com", "toqxzbprjigwaqib")

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	fmt.Println("Email sent successfully! Verification code: " + verificationCode)

	return nil
}
