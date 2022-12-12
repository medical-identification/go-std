package middleware

import (
	"fmt"
	"os"

	"github.com/medical-identification/go-std/helper"
	"github.com/medical-identification/go-std/model"

	"github.com/medical-identification/go-std/http"

	"github.com/gofiber/fiber/v2"
)

func Authorization() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authorization := ctx.Get("Authorization")
		// checking if the authorization header is empty.
		if authorization == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
				Message: "Authorization token missing",
			})
		}

		profileUrl := fmt.Sprintf("%v/auth/profile", os.Getenv("AUTH_SERVER"))

		var user model.UserWithRelations
		err := http.GetFromJson(profileUrl, authorization, &user)
		if err != nil {
			return helper.UnauthorizedError(err, ctx)
		}

		if user == (model.UserWithRelations{}) {
			return helper.UnauthorizedError(fmt.Errorf("missing or malfunction token"), ctx)
		}

		// passing user to the request scope.
		ctx.Locals("user", user)

		return ctx.Next()
	}
}
