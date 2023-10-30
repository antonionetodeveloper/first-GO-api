package handlers

import (
	"api/models"
	"api/models/operations"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		log.Printf("Error while parsing id: %v", err)
		http.Error(writer,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}

	var user models.User

	err = json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writer,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}

	rows, err := operations.UpdateUserByID(int64(id), user)
	if err != nil {
		log.Printf("Error while updating user: %v", err)
		http.Error(writer,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}

	if rows > 1 {
		log.Printf("Error: %d field were updated.", rows)
	}

	response := map[string]any{
		"Success": true,
		"Message": fmt.Sprintf("User successfuly updated! User id: %d", id),
	}

	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
