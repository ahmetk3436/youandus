package router

import (
	"github.com/gofiber/fiber/v2"
	"youandus/pkg/users/handler"
	"youandus/pkg/users/middleware"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Router struct {
	Handler         handler.Handler
	WebApiFramework *fiber.App
}

func NewRouter(handler handler.Handler) *Router {
	return &Router{
		Handler: handler,
	}
}
func (r *Router) InitRouter() *Router {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "*",
	}))
	login := app.Group("")
	login.Post("/register", r.Handler.Api.CreateUser)
	login.Post("/login", r.Handler.Api.LoginUser)
	users := app.Group("users")
	users.Use(middleware.Auth)
	users.Delete("", r.Handler.Api.DeleteUser)
	users.Put("", r.Handler.Api.UpdateUser)
	//users.Get("", r.Handler.Api.GetUsers)

	r.WebApiFramework = app
	return r
}
