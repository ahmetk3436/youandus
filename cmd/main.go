package main

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"sync"
	event "youandus/pkg/event/handler"
	info "youandus/pkg/info/handler"
	"youandus/pkg/info/repository"
	profileHandler "youandus/pkg/profile/handler"
	"youandus/pkg/users/handler"
	"youandus/pkg/users/router"
	"youandus/pkg/utility/api"
)

func main() {
	startServer()
}

func startServer() {
	var wg sync.WaitGroup
	wg.Add(1)
	stockHandlerApi := handler.InitUsers()
	e := router.NewRouter(*stockHandlerApi).InitRouter()
	ServeApi(e.WebApiFramework)
	go EmailConsumer(&wg)
	ProfileHandler(e.WebApiFramework)
	EventHandler(e.WebApiFramework)
	e.WebApiFramework.Get("/location", func(c *fiber.Ctx) error {
		ip := c.IP()
		data, err := api.GetLocationInfo(ip)
		if err != nil {
			c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"Message": "IP bilgileri alınamadı : " + err.Error(),
			})
		}
		return c.Status(200).JSON(data)
	})
	err := e.WebApiFramework.Listen(":1323")
	if err != nil {
		panic(err)
	}
}
func ServeApi(app *fiber.App) {
	info.Init(app)
}

func EmailConsumer(wg *sync.WaitGroup) {
	defer wg.Done()
	repository.ConsumeVerification()
}

func ProfileHandler(fiber *fiber.App) {
	profileHandler.InitProfile(fiber)
}

func EventHandler(fiber *fiber.App) {
	event.InitEvent(fiber)
}
