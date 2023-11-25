package helper

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetUserID(c *fiber.Ctx) (uint, error) {
	user, ok := c.Locals("user_id").(string)
	if !ok || user == "" {
		return 0, errors.New("user_id not found in token")
	}
	userID, err := strconv.Atoi(user)
	if err != nil {
		return 0, errors.New("user id dönüştürülemedi")
	}
	return uint(userID), nil
}

func GetUserName(c *fiber.Ctx) (string, error) {
	userName := c.Get("userName")
	if userName == "" {
		return "", errors.New("user_name not found in token")
	}
	return userName, nil
}

type JwtCustom struct {
	UserID   string `json:"user_id"`
	UserName string `json:"userName"`
	jwt.StandardClaims
}

func Auth(ctx *fiber.Ctx) error {
	token, err := getTokenFromHeader(ctx)
	if err != nil {
		return ctx.Status(403).JSON(fiber.Map{"message": "Oturum bulunamadı. " + err.Error()})
	}

	ctx.Locals("user_id", token.UserID)
	ctx.Locals("user_name", token.UserName)
	return ctx.Next()
}

func getTokenFromHeader(c *fiber.Ctx) (*JwtCustom, error) {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return &JwtCustom{}, errors.New("authorization header not found")
	}

	// Token başlığını ayrıştır
	tokenString := strings.Split(authHeader, "Bearer ")
	if len(tokenString) != 2 {
		return &JwtCustom{}, errors.New("invalid token format")
	}

	token, err := jwt.ParseWithClaims(tokenString[1], &JwtCustom{}, func(token *jwt.Token) (interface{}, error) {
		// Verify the token signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte("youanduseventplannerv1"), nil
	})

	if err != nil {
		return &JwtCustom{}, err
	}

	if !token.Valid {
		return &JwtCustom{}, errors.New("invalid token")
	}

	jwtCustomClaims, ok := token.Claims.(*JwtCustom)
	if !ok {
		return &JwtCustom{}, errors.New("invalid claims")
	}

	return jwtCustomClaims, nil
}
