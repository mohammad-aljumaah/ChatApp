package handlers

import (
	"net/http"

	"github.com/mohammad-aljumaah/ChatApp/auth/internal/service"
)

type Handler struct {
	Service *service.Service
}

func NewHandler() *Handler {
	return &Handler{
		Service: service.NewService(),
	}
}

type RegisterRequest struct {
	Email    string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	err := ReadJSON(w, r, &req)
	if err != nil {
		ErrorJSON(w, err, http.StatusBadRequest)
	}

	err = h.Service.Register(req.Email, req.Password)
	if err != nil {
		ErrorJSON(w, err, http.StatusBadRequest)
	}

	var payload JSONResponse
	payload.Error = false
	payload.Message = "User registered successfully"

	WriteJSON(w, http.StatusOK, payload)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if err := WriteJSON(w, http.StatusOK, "Hello World"); err != nil {
		panic(err)
	}
}
