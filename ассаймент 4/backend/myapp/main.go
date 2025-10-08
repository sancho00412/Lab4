package main

import (
	"log"
	"net/http"

	"myapp/handlers" // Импортируем пакет с обработчиками

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Обработчик для логина
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	// Запуск HTTP-сервера
	log.Fatal(http.ListenAndServe(":8080", r))
}
