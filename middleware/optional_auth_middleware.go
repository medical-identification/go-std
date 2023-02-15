package middleware

import (
	"fmt"
	"os"

	"github.com/medical-identification/go-std/model"

	"github.com/medical-identification/go-std/http"

	"github.com/gofiber/fiber/v2"
)

func OptionalAuthorization() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authorization := ctx.Get("Authorization")
		// checking if the authorization header is empty.
		if authorization == "" {
			ctx.Next()
		}

		profileUrl := fmt.Sprintf("%v/auth/profile", os.Getenv("AUTH_SERVER"))

		var user model.UserWithRelations
		err := http.GetFromJson(http.MethodGet, profileUrl, authorization, &user, nil)
		if err != nil {
			ctx.Next()
		}

		if user == (model.UserWithRelations{}) {
			ctx.Next()
		}
		// passing user to the request scope.
		ctx.Locals("user", user)

		return ctx.Next()
	}
}
