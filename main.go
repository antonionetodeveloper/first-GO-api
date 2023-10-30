package main

import (
	config "api/configs"
	"api/handlers"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	err := config.Load()
	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()
	router.Get("/user/{id}", handlers.GetUserByID)
	router.Post("/user/create", handlers.CreateUser)
	router.Put("/user/edit/{id}", handlers.UpdateUser)
	router.Delete("/user/delete/{id}", handlers.DeleteUser)

	http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), router)
}
