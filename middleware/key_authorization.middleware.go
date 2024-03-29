package middleware

import (
	"fmt"
	"os"

	"github.com/medical-identification/go-std/model"
	"github.com/medical-identification/go-std/util"

	"github.com/gofiber/fiber/v2"
	"github.com/medical-identification/go-std/http"
)

func ApiKeyAuthorization() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		WHITE_LIST_KEY := ctx.Get("WHITE-LIST-KEY")
		if WHITE_LIST_KEY == os.Getenv("WHITE_LIST_KEY") || util.ParseBoolean(os.Getenv("IGNORE_KEY")) {
			return ctx.Next()
		}

		ACCESS_KEY := ctx.Get("ACCESS-KEY")
		errMsg := "Invalid or missing ACCESS-KEY"

		if ACCESS_KEY == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
				Message: errMsg,
			})
		}

		url := fmt.Sprintf("%v/key/%v", os.Getenv("MID_AUTH_SERVER"), ctx.Get("ACCESS-KEY"))

		var body interface{}
		err := http.GetAnonymousFromJson(http.MethodGet, url, &body, nil)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
				Message: errMsg,
			})
		}

		// adding api key to the request key that is affected by the middleware
		ctx.Locals("ACCESS-KEY", ACCESS_KEY)

		return ctx.Next()
	}
}
