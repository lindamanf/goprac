package services

import (
	"work/02_mvc/domain"
	"work/02_mvc/utils"
)

func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
