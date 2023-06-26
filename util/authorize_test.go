package util

import (
	"fmt"
	"testing"
)

func TestExtractUser(t *testing.T) {
	t.Setenv("MID_AUTH_SERVER", "http://localhost:5000/api")
	// this is a demo token
	user, err := ExtractUserFromToken("Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJtaWQiOiJNSURYQkcyU0FVOCIsInNlc3Npb24iOiIyZGY5RkRHVTNmNk5falBzYTI2SDQiLCJ0aW1lIjoxNjg1OTY2MTYzfQ.73yj244Pe3L2eu2yMQqtkcrf2CPb4kjt0TDNxfnv0Yo")

	if err != nil {
		t.Fatal("Error occurred", err)
	}

	fmt.Println("user", user)
}
