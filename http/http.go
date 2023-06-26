// mid http package
package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

// Common HTTP methods.
//
// Unless otherwise noted, these are defined in RFC 7231 section 4.3.
const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH" // RFC 5789
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)

var client *http.Client = &http.Client{}

// GetFromJson send http request to mid auth server
func GetFromJson(method, url string, authorization string, target interface{}, requestBody io.Reader) error {
	// initializing a new request to the server
	req, err := http.NewRequest(method, url, requestBody)
	// if error encounter , throw it
	if err != nil {
		return err
	}

	// some http end points in mid server requires to be authenticated
	// we need to add the authorization header which the server uses
	// so we send the user token which is fetches from client device (browser/web).
	req.Header.Set("Authorization", authorization)
	req.Header.Set("WHITE-LIST-KEY", os.Getenv("WHITE_LIST_KEY"))

	// finally, send the request to the server
	resp, err := client.Do(req)
	// if an error is encountered, return the response error
	if err != nil {
		fmt.Println("Error when sending request to the server")
		return err
	}

	// finally, very important, close the response body las las.
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// fmt.Printf("Error %s", err)
		return err
	}

	if resp.StatusCode != fiber.StatusOK {
		// decoding body
		var errResponse interface{}
		err := json.Unmarshal([]byte(body), &errResponse)
		if err != nil {
			return err
		}

		switch errResponse.(type) {
		case map[string]interface{}:
			e := errResponse.(map[string]interface{})
			return errors.New(e["message"].(string))
		default:
			return fmt.Errorf(string(body))
		}
	}

	// decode this new data and parse it to the target whic h is a memory pointer
	// the decoder package receives the object with it memory address and modifies it
	// directly
	// return json.NewDecoder(resp.Body).Decode(target) // !!! solution not working
	return json.Unmarshal(body, target)
}

func GetAnonymousFromJson(method, url string, target interface{}, requestBody io.Reader) error {
	// initializing a new request to the server
	req, err := http.NewRequest(method, url, requestBody)
	// if error encounter , throw it
	if err != nil {
		return err
	}

	// headers
	req.Header.Set("WHITE-LIST-KEY", os.Getenv("WHITE_LIST_KEY"))

	// finally, send the request to the server
	resp, err := client.Do(req)
	// if an error is encountered, return the response error
	if err != nil {
		fmt.Println("Error when sending request to the server")
		return err
	}

	// finally, very important, close the response body las las
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error %s", err)
		return err
	}

	if resp.StatusCode != fiber.StatusOK {
		// unmarshal body
		var errResponse interface{}
		err := json.Unmarshal([]byte(body), &errResponse)
		if err != nil {
			return err
		}

		switch errResponse.(type) {
		case map[string]interface{}:
			e := errResponse.(map[string]interface{})
			return errors.New(e["message"].(string))
		default:
			return fmt.Errorf(string(body))
		}
	}

	// fmt.Printf("Body : %s\n", body)

	// decode this new data and parse it to the target which is a memory pointer
	// the decoder package receives the object with it memory address and modifies it
	// directly
	// return json.NewDecoder(resp.Body).Decode(target)
	return json.Unmarshal(body, target)
}
