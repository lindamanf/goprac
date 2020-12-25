package domain

import (
	"fmt"
	"log"
	"net/http"
	"work/mvc/utils"
)

var (
	users = map[int64]*User{
		123: &User{ID: 123, FirstName: "Fumiya", LastName: "Hayashi", Email: "fhayashi843@gmail.com"},
	}

	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {
	log.Println("we're accessing the database")
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user with id %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
