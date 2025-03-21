package handler

import (
	"analytics_project/internal/service"
	"analytics_project/pkg/response"
	"net/http"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) GetAllUsersWithMenu(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.GetAllUsersWithMenu()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch users with menu")
		return
	}
	response.JSON(w, http.StatusOK, "Users with menu fetched successfully", users)
}

func (h *UserHandler) RunQuery(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		response.Error(w, http.StatusBadRequest, "Name parameter is required")
		return
	}

	user, err := h.Service.RunServiceQuery(name)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch user")
		return
	}

	if user == nil {
		response.Error(w, http.StatusNotFound, "User not found")
		return
	}

	response.JSON(w, http.StatusOK, "User fetched successfully", user)
}
