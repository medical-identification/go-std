package middleware

import (
	"github.com/medical-identification/go-std/helper"
	"github.com/medical-identification/go-std/util"

	"github.com/gofiber/fiber/v2"
)

func Authorization() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user, err := util.ExtractUserFromToken(ctx.Get("Authorization"))
		if err != nil {
			helper.UnauthorizedError(err, ctx)
		}

		// passing user to the request scope.
		ctx.Locals("user", user)

		return ctx.Next()
	}
}
