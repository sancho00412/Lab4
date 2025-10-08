package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Временная in-memory "база" пользователей (только для теста).
// Хэш для пароля "password123" с cost=12.
// В реальном приложении замените на запрос к БД.
var users = map[string]string{
	"john": "$2b$12$dNzMIobQ1SVTaM2HixrNqOaUUPv3fX55NsnKUi1qc6jZq20wt0w6C",
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token,omitempty"`
	Error string `json:"error,omitempty"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Защита от больших тел
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1 MiB

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var req loginRequest
	if err := dec.Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(loginResponse{Error: "invalid request body"})
		return
	}

	req.Username = strings.TrimSpace(req.Username)
	if req.Username == "" || req.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(loginResponse{Error: "username and password required"})
		return
	}

	storedHash, ok := users[req.Username]
	// Унифицированный ответ при неверных учётных данных
	if !ok || !CheckPasswordHash(req.Password, storedHash) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(loginResponse{Error: "invalid credentials"})
		return
	}

	token, err := GenerateJWT(req.Username)
	if err != nil {
		// Логирование ошибки на сервере (не отдаем клиенту детали)
		// log.Printf("GenerateJWT error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(loginResponse{Error: "internal server error"})
		return
	}

	json.NewEncoder(w).Encode(loginResponse{Token: token})
}
