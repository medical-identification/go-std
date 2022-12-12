package http

import (
	"testing"
)

func TestError(t *testing.T) {
	url := "http://localhost:5000/user/identity/aaa"
	var body interface{}
	err := GetAnonymousFromJson(url, &body)

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
