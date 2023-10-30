package handlers

import (
	"api/models"
	"api/models/operations"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func GetUserByID(writer http.ResponseWriter, request *http.Request) {
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
		log.Printf("Error while decoding json: %v", err)
		http.Error(writer,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}

	log.Println(int64(id))
	user, err = operations.GetUserByID(int64(id))
	if err != nil {
		log.Printf("Error while searching user: %v", err)
		http.Error(writer,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}

	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(user)
}
