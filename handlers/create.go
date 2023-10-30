package handlers

import (
	"api/models"
	"api/models/operations"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUser(writer http.ResponseWriter, request *http.Request) {
	var user models.User

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writer,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}

	var response map[string]any
	id, err := operations.CreateUser(user)
	if err != nil {
		response = map[string]any{
			"Success": false,
			"Message": fmt.Sprintf("An error ocurred while trying to instert user: %v", err),
		}
	} else {
		response = map[string]any{
			"Success": true,
			"Message": fmt.Sprintf("User successfuly added! User id: %d", id),
		}
	}

	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
