package handlers

import (
	"api/models/operations"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		log.Printf("Error while parsing id: %v", err)
		http.Error(writer,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}

	_, err = operations.DeleteUserByID(int64(id))
	if err != nil {
		log.Printf("Error while deleting user: %v", err)
		http.Error(writer,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}

	response := map[string]any{
		"Success": true,
		"Message": fmt.Sprintf("User successfuly deleted!"),
	}

	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
