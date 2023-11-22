package handler

import (
	"time"
	"youandus/internal/storage"
	"youandus/pkg/info/api"
	"youandus/pkg/info/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Init(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "*",
	}))
	app.Static("/verification", "/infoapiservice/pkg/handler/html")
	app.Get("/verification/dogrulama-basarili.html", func(c *fiber.Ctx) error {
		time.Sleep(time.Second * 2)
		c.Redirect("https://youandus.net/signin.html")
		return nil
	})
	db := storage.GetDB()
	repo := repository.NewVerificationRepo(db)
	api := api.NewVerificationAPI(app, repo)
	app.Get("/api/verifyemail", api.HandleVerification)
}
