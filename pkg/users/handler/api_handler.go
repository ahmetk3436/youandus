package handler

import (
	"fmt"
	"youandus/pkg/users/api"
	"youandus/pkg/users/repository"
	"youandus/pkg/users/service"

	"youandus/internal/storage"
)

type Handler struct {
	Api *api.UsersAPI
}

func NewUsersHandler() *Handler {
	dbInstance := storage.GetDB()
	ch := storage.ConnectRabbitMQ()
	redisInstance, err := storage.NewRedisClient("localhost:6379", "toor")
	if err != nil {
		fmt.Println(err.Error())
	}
	//user
	userRepository := repository.NewUsersRepository(dbInstance, redisInstance, ch)
	userService := service.NewUsersService(userRepository)
	userApi := api.NewUsersAPI(userService)
	return &Handler{
		Api: userApi}
}
