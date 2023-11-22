package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
	"github.com/streadway/amqp"
)

func ConsumeVerification(ch *amqp.Channel) {
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
	// Mesajları işle
	go func() {
		for msg := range msgs {
			// Mesajı doğrulama işlemine tabi tut
			var user User
			err := json.Unmarshal(msg.Body, &user)
			if err != nil {
				log.Fatalf("Could not decode JSON: %s", err)
			}

			// Doğrulama işlemini gerçekleştir
			SendEmail(user.ID, user.Email)

			fmt.Printf("Verification code received: %d\n", user.ID)
		}
	}()
}

type User struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

func ConsumeNotification(ch *amqp.Channel) {
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

func SendEmail(userID uint, userEmail string) {
	// Kullanıcının ID'sini kullanarak kullanıcıyı veritabanından alın veya gerekli işlemleri gerçekleştirin
	// Örneğin:
	// user := getUserByID(userID)
	// ...

	// E-posta gönderme işlemleri
	e := email.NewEmail()
	e.From = "testerahmet16@gmail.com"
	e.To = []string{"bekirbayar2@gmail.com"}
	e.Subject = "Verification Code"
	e.Text = []byte("Please click the link and verificate your account !")

	err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "testerahmet16@gmail.com", "123456test", "smtp.gmail.com"))
	if err != nil {
		log.Fatalf("Could not send email: %s", err)
	}

	fmt.Println("Email sent successfully!")
}
