// mid http package
package http

import (
	"net/http"
	"testing"
)

func TestError(t *testing.T) {
	url := "http://localhost:5000/api/auth/profile"
	var body interface{}
	// err := GetAnonymousFromJson(url, &body)
	// this is a test token and remember it invalid
	err := GetFromJson(http.MethodGet, url, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJtaWQiOiJNSURCQUVQN0RPSiIsInNlc3Npb24iOiJhMEtoSmxzVTZfMDAyVnVMZWlKQkcifQ.3gkzURkgYU0LxyvQtKGEGNcph-JC7Ud5nYttWbHrfS", &body, nil)

	if err != nil {
		// var errResponse interface{}
		// err := json.Unmarshal([]byte(err.Error()), &errResponse)
		// if err != nil {
		// 	t.Error("error", err)
		// }
		// e := errResponse.(map[string]interface{})
		// t.Error("Error occurred", e["message"].(string))
		t.Error("Error occurred", err)
	}
}
