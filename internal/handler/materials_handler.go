package handler

import (
	"context"
	"go-minimal/internal/service"
	"go-minimal/internal/utils"
	"net/http"
)

type MaterialHandler struct {
	service *service.MaterialsService
}

func NewMaterialHandler(service *service.MaterialsService) *MaterialHandler {
	if service == nil {
		panic("Material service cannot be nil")
	}

	return &MaterialHandler{
		service: service,
	}
}

func (h *MaterialHandler) GetAllMaterial(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Invalid Method", "Method Not Allowed")
	}

	materials, err := h.service.GetAllMaterial(context.Background())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error(), "Internal Server Error")
	}

	w.Header().Set("Content-Type", "application/json")
	utils.WriteSuccess(w, http.StatusOK, "materials fetched successfully", materials)
}
