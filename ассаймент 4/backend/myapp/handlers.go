package handlers

import (
	"fmt"
	"net/http"
)

// LoginHandler - обработчик для маршрута /login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Login Handler")
}
