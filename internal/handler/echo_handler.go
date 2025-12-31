package handler

import (
	"encoding/json"
	"net/http"

	"goc/internal/domain"
	"goc/internal/service"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	var req domain.EchoRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	resp := service.Echo(req.Message)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
