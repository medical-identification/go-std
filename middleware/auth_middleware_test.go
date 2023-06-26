package middleware

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJtaWQiOiJNSURYQkcyU0FVOCIsInNlc3Npb24iOiIyZGY5RkRHVTNmNk5falBzYTI2SDQiLCJ0aW1lIjoxNjg1OTY2MTYzfQ.73yj244Pe3L2eu2yMQqtkcrf2CPb4kjt0TDNxfnv0Yo"

func TestAuthMiddleware(t *testing.T) {
	t.Setenv("MID_AUTH_SERVER", "http://localhost:5000/api")
	app := fiber.New()
	app.Use(Authorization())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal("error occurred: ", err)
	}
	if body, err := ioutil.ReadAll(resp.Body); err == nil {
		log.Println(string(body)) // "testing"
	}
}
