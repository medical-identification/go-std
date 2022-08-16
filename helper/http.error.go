package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/medical-identification/go-std/model"
)

func UnauthorizedError(err error, ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
		Message: err.Error(),
	})
}
