package services

import (
	"work/mvc/domain"
	"work/mvc/utils"
)

func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
