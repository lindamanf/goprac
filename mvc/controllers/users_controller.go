package controllers

import (
	"net/http"
	"strconv"
	"work/mvc/services"
	"work/mvc/utils"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userID, err := (strconv.ParseInt(c.Param("user_id"), 10, 64))
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	user, apiErr := services.UsersService.GetUser(userID)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, user)
}
