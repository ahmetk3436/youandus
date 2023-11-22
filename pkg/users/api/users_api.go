package api

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"youandus/helper"
	"youandus/pkg/users/model/user"
	"youandus/pkg/users/service"
)

type UsersAPI struct {
	usersService *service.UsersService
}

func NewUsersAPI(service *service.UsersService) *UsersAPI {
	return &UsersAPI{
		usersService: service,
	}
}

func (api *UsersAPI) CreateUser(c *fiber.Ctx) error {
	user := new(user.UserRegister)
	if err := c.BodyParser(user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	createdUser, err := api.usersService.CreateUser(user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(createdUser)
}
func (api *UsersAPI) LoginUser(c *fiber.Ctx) error {
	user := new(user.UserLogin)
	if err := c.BodyParser(user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	loginUser, err := api.usersService.LoginUser(user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(loginUser)
}
func (api *UsersAPI) UpdateUser(c *fiber.Ctx) error {
	userID, _ := helper.GetUserID(c)
	user := new(user.UserRegister)
	if err := c.BodyParser(user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}
	updatedUser, err := api.usersService.UpdateUser(userID, user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}
	return c.Status(http.StatusOK).JSON(updatedUser)
}
func (api *UsersAPI) DeleteUser(c *fiber.Ctx) error {
	id, _ := helper.GetUserID(c)
	if err := api.usersService.DeleteUser(strconv.Itoa(int(id))); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User deleted",
	})
}

func (api *UsersAPI) GetUsers(c *fiber.Ctx) error {
	users, err := api.usersService.GetUsers()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}
	return c.Status(http.StatusOK).JSON(users)
}
