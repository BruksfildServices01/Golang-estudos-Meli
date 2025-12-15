package apicriacaousuarios

import (
	"encoding/json"
	"net/http"
)

type Server struct {
	userService *UserService
}

func NewServe() *Server {
	return &Server{
		userService: NovoService(),
	}
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
