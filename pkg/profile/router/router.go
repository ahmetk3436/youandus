package router

import (
	"youanduseventplanner/profile_service/pkg/api"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	api   *api.Api
	fiber *fiber.App
}

func NewRouter(api *api.Api, fiber *fiber.App) Router {
	return Router{
		api:   api,
		fiber: fiber,
	}
}

func (r Router) InitRouter() {
	r.fiber.Get("/profile", r.api.GetProfile)
	r.fiber.Post("/profile", r.api.CreateProfile)
	r.fiber.Put("/profile", r.api.UpdateProfile)
	r.fiber.Delete("/profile", r.api.DeleteProfile)
}
