package helper

import (
	"github.com/gofiber/fiber/v2"
)

func UnauthorizedError(body interface{}, ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(body)
}
