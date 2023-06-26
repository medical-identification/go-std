package util

import (
	"errors"
	"fmt"
	"os"

	"github.com/medical-identification/go-std/http"
	"github.com/medical-identification/go-std/model"
)

func ExtractUserFromToken(authorization string) (model.UserWithRelations, error) {
	// checking if the authorization header is empty.
	if authorization == "" {
		return model.UserWithRelations{}, errors.New("Authorization token missing")
	}

	profileUrl := fmt.Sprintf("%v/auth/profile", os.Getenv("MID_AUTH_SERVER"))

	var user model.UserWithRelations
	err := http.GetFromJson(http.MethodGet, profileUrl, authorization, &user, nil)
	if err != nil {
		return model.UserWithRelations{}, err
	}

	if user == (model.UserWithRelations{}) {
		return model.UserWithRelations{}, errors.New("Missing or malfunction token")
	}

	return user, nil
}
