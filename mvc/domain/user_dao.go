package domain

import (
	"fmt"
	"net/http"
	"work/mvc/utils"
)

var (
	users = map[int64]*User{
		123: &User{ID: 123, FirstName: "Fumiya", LastName: "Hayashi", Email: "fhayashi843@gmail.com"},
	}
)

func GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user with id %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
