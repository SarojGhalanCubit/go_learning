package colorHandler

import (
	"context"
	"encoding/json"
	colorModel "go-minimal/internal/modules/colors/model"
	colorService "go-minimal/internal/modules/colors/service"
	"go-minimal/internal/utils"
	validateColor "go-minimal/internal/utils/validate"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ColorHandler struct {
	service *colorService.ColorsService
}

func NewColorHandler(service *colorService.ColorsService) *ColorHandler {
	return &ColorHandler{service: service}
}

func (h *ColorHandler) GetAllColors(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Invalid Method", "Method Not Allowed")
		return
	}

	colors, err := h.service.GetAllColors(context.Background())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error(), "Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.WriteSuccess(w, http.StatusOK, "colors fetched successfully", colors)
}

func (h *ColorHandler) CreateColor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Invalid method", "Method not allowed")
		return
	}

	var color colorModel.CreateColor
	err := json.NewDecoder(r.Body).Decode(&color)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "create material failed", "Invalid request body")
		return
	}

	if userValidationErr := validateColor.ValidateColor(color.Name, color.HexCode); len(userValidationErr) > 0 {
		utils.WriteError(w, http.StatusUnprocessableEntity, "Validation Error", userValidationErr)
		return
	}

	createdMaterial, err := h.service.CreateColor(context.Background(), color)

	if err != nil {

		utils.WriteError(
			w,
			http.StatusInternalServerError,
			"Internal server error",
			nil,
		)

		return
	}

	utils.WriteSuccess(w, http.StatusOK, "color created successfully", createdMaterial)
}

func (h *ColorHandler) GeyByColorID(w http.ResponseWriter, r *http.Request) {

	IDstr := chi.URLParam(r, "id")

	if r.Method != http.MethodGet {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Invalid Method", "Method Not Allowed")
		return
	}

	user, err := h.service.GeyByColorID(context.Background(), IDstr)

	if err != nil {
		utils.WriteError(w, http.StatusNotFound, "Request Failed", err.Error())
		return
	}

	utils.WriteSuccess(w, http.StatusOK, "color fetched successfully", user)

}

func (h *ColorHandler) UpdateColor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Invalid Method", "Method Not Allowed")
		return
	}
	var color colorModel.CreateColor

	if err := json.NewDecoder(r.Body).Decode(&color); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Update material failed", "Invalid Request Body")
		return
	}
	if userValidationErr := validateColor.ValidateColor(color.Name, color.HexCode); len(userValidationErr) > 0 {
		utils.WriteError(w, http.StatusUnprocessableEntity, "Validation Error", userValidationErr)
		return
	}

	IDstr := chi.URLParam(r, "id")

	updated, err := h.service.UpdateColor(context.Background(), IDstr, color)
	if err != nil {
		if err.Error() == "requested color did not exist" {
			utils.WriteError(w, http.StatusConflict, err.Error(), "color update failed")
			return
		}

		utils.WriteError(
			w,
			http.StatusInternalServerError,
			"Internal server error",
			nil,
		)

		return
	}
	utils.WriteSuccess(w, http.StatusCreated, "color updated successfully", updated)
}

func (h *ColorHandler) DeleteColorsByID(w http.ResponseWriter, r *http.Request) {

	IDstr := chi.URLParam(r, "id")
	if r.Method != http.MethodDelete {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Invalid Method", "Method Not Allowed")
		return
	}

	deletedColor, err := h.service.DeleteColorsByID(context.Background(), IDstr)

	if err != nil {
		utils.WriteError(w, http.StatusNotFound, "Request Failed", err.Error())
		return
	}

	utils.WriteSuccess(w, http.StatusOK, "color deleted successfully", deletedColor)

}
