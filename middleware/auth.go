package middleware

import (
	"butter/exception"
	"butter/feature/user/service"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthMiddleware struct {
	UserService service.UserService
	DB          *gorm.DB
}

func NewAuthMiddleware(userService service.UserService, db *gorm.DB) *AuthMiddleware {
	return &AuthMiddleware{
		UserService: userService,
		DB:          db,
	}
}

func (a *AuthMiddleware) AuthenticateFiber(allowedGuest bool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var tokenString string

		spliten := strings.Split(c.Get("authorization"), "Bearer ")
		if len(spliten) == 2 {
			tokenString = spliten[1]
		}

		if allowedGuest && tokenString == "" {
			return c.Next()
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			return exception.NewBadRequestError("unauthorized")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				return exception.NewBadRequestError("unauthorized")
			}

			id, ok := claims["sub"].(string)
			if !ok {
				return exception.NewBadRequestError("unauthorized")
			}

			user := a.UserService.FindById(a.DB, id)

			if user.Id == "" {
				return exception.NewBadRequestError("unauthorized")
			}

			c.Locals("user_id", user.Id)
			c.Locals("email", user.Email)
			c.Locals("username", user.Username)

			return c.Next()
		} else {
			return exception.NewBadRequestError("unauthorized")
		}
	}
}
