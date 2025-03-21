package handler

import (
	"analytics_project/internal/util"
	"analytics_project/pkg/response"
	"encoding/json"
	"net/http"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

type TokenRequest struct {
	UserName string `json:"user_name"`
}

func (h *AuthHandler) CreateToken(w http.ResponseWriter, r *http.Request) {
	var req TokenRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.UserName == "" {
		response.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	token, err := util.GenerateToken(req.UserName)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	response.JSON(w, http.StatusOK, "Token generated successfully", map[string]string{"token": token})
}
