package middleware

// import (
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"testing"

// 	"github.com/gofiber/fiber/v2"
// )

// func TestAdd(t *testing.T) {
// 	app := fiber.New()
// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.JSON(fiber.Map{
// 			"message": "Hello, World!",
// 		})
// 	})

// 	req, _ := http.NewRequest(http.MethodGet, "/", nil)
// 	// req.Header.Set("Authorization", "")
// 	resp, err := app.Test(req)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if body, err := ioutil.ReadAll(resp.Body); err == nil {
// 		log.Println(string(body)) // "testing"
// 	}
// }
