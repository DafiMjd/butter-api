package middleware

import (
	"butter/helper"
	"butter/pkg/exception"
	"butter/pkg/user"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthMiddleware struct {
	UserService user.UserService
	DB          *gorm.DB
}

func NewAuthMiddleware(userService user.UserService, db *gorm.DB) *AuthMiddleware {
	return &AuthMiddleware{
		UserService: userService,
		DB:          db,
	}
}

func (a *AuthMiddleware) AuthenticateFiber(allowedGuest bool) fiber.Handler {
	checkToken := checkToken

	return func(c *fiber.Ctx) error {
		var tokenString string

		spliten := strings.Split(c.Get("authorization"), "Bearer ")
		if len(spliten) == 2 {
			tokenString = spliten[1]
		}

		if allowedGuest && tokenString == "" {
			return c.Next()
		}

		return checkToken(a, c, tokenString)
	}
}

func (a *AuthMiddleware) AuthenticateRefreshToken() fiber.Handler {
	checkToken := checkToken

	return func(c *fiber.Ctx) error {

		payload := struct {
			RefreshToken string `json:"refreshToken"`
		}{}

		err := c.BodyParser(&payload)
		if err != nil {
			panic(exception.NewBadRequestError(err.Error()))
		}

		return checkToken(a, c, payload.RefreshToken)
	}
}

func checkToken(
	a *AuthMiddleware,
	c *fiber.Ctx,
	t string,
) error {
	token, err := helper.ParseJwt(t)

	if err != nil {
		return exception.NewBadRequestError(err.Error())
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return exception.NewUnauthenticatedError("token expired")
		}

		id, ok := claims["sub"].(string)
		if !ok {
			return exception.NewUnauthenticatedError("unauthorized")
		}

		fmt.Println("disini")

		c.Locals("user_id", id)

		return c.Next()
	} else {
		return exception.NewUnauthenticatedError("unauthorized")
	}
}
