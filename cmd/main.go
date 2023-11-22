package main

import (
	"github.com/gofiber/fiber/v2"
	"sync"
	info "youandus/pkg/info/handler"
	"youandus/pkg/info/repository"
	profileHandler "youandus/pkg/profile/handler"
	"youandus/pkg/users/handler"
	"youandus/pkg/users/router"
)

func main() {
	startServer()
}

func startServer() {
	var wg sync.WaitGroup
	wg.Add(1)
	stockHandlerApi := handler.NewUsersHandler()
	e := router.NewRouter(*stockHandlerApi).InitRouter()
	ServeApi(e.WebApiFramework)
	go EmailConsumer(&wg)
	ProfileHandler(e.WebApiFramework)
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
	profileHandler.InitHandler(fiber)
}
