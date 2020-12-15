package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"work/mvc/services"
	"work/mvc/utils"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userIDParam := req.URL.Query().Get("user_id")
	userID, err := (strconv.ParseInt(userIDParam, 10, 64))
	if err != nil {
		// Just return the Bad Request to the client
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(http.StatusNotFound)
		res.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userID)
	if apiErr != nil {
		// Handle the err and return to the client
		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write(jsonValue)
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}
