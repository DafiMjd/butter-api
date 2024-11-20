package exception

import (
	"butter/pkg/model"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {

	var o Oops

	if errors.As(err, &o) {
		return c.Status(o.code).JSON(model.WebResponse{
			Code:    o.code,
			Status:  o.status,
			Message: o.message,
			Data:    o.data,
		})
	}

	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return c.Status(code).JSON(model.WebResponse{
		Code:    http.StatusInternalServerError,
		Status:  "INTERNAL_SERVER_ERROR",
		Message: err.Error(),
		Data:    o.data,
	})

}
