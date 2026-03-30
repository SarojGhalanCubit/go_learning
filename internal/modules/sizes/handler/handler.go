package sizeHandler

import (
	"context"
	"encoding/json"
	sizeModel "go-minimal/internal/modules/sizes/model"
	sizeService "go-minimal/internal/modules/sizes/service"
	"go-minimal/internal/utils"

	validateaSize "go-minimal/internal/utils/validate"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type SizeHandler struct {
	service *sizeService.SizeService
}

func NewSizeHandler(s *sizeService.SizeService) *SizeHandler {
	return &SizeHandler{service: s}
}

func (s *SizeHandler) GetAllSizes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Invalid request method ", "method not allowed")
	}

	sizes, err := s.service.GetAllSizes(context.Background())

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error(), "Internal Server Error")
	}

	w.Header().Set("Content-Type", "application/json")
	utils.WriteSuccess(w, http.StatusOK, "sizes fetched successfully", sizes)

}

func (s *SizeHandler) CreateSize(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteError(w, http.StatusMethodNotAllowed, "invalid request method", "method not allowed")
	}

	var size sizeModel.CreateSize
	err := json.NewDecoder(r.Body).Decode(&size)
	if err != nil {

		utils.WriteError(w, http.StatusInternalServerError, "create material failed", "Invalid request body")
	}

	if userValidationErr := validateaSize.ValidateSize(size.Name, size.SortOrder); len(userValidationErr) > 0 {
		utils.WriteError(w, http.StatusUnprocessableEntity, "Validation Error", userValidationErr)
		return
	}

	createdMaterial, err := s.service.CreateSize(context.Background(), size)
	if err != nil {

		utils.WriteError(
			w,
			http.StatusInternalServerError,
			"Internal server error",
			nil,
		)

		return
	}

	utils.WriteSuccess(w, http.StatusOK, "size created successfully", createdMaterial)

}

func (h *SizeHandler) UpdateSize(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Invalid Method", "Method Not Allowed")
		return
	}

	var size sizeModel.CreateSize

	if err := json.NewDecoder(r.Body).Decode(&size); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Update size failed", "Invalid Request Body")
		return
	}

	if userValidationErr := validateaSize.ValidateSize(size.Name, size.SortOrder); len(userValidationErr) > 0 {
		utils.WriteError(w, http.StatusUnprocessableEntity, "Validation Error", userValidationErr)
		return
	}

	IDstr := chi.URLParam(r, "id")

	updated, err := h.service.UpdateSize(context.Background(), IDstr, size)
	if err != nil {
		if err.Error() == "requested size did not exist" {
			utils.WriteError(w, http.StatusConflict, err.Error(), "size update failed")
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
	utils.WriteSuccess(w, http.StatusCreated, "size updated successfully", updated)
}

func (h *SizeHandler) DeleteSizeByID(w http.ResponseWriter, r *http.Request) {

	IDstr := chi.URLParam(r, "id")
	if r.Method != http.MethodDelete {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Invalid Method", "Method Not Allowed")
		return
	}

	deletedMaterial, err := h.service.DeleteSizeByID(context.Background(), IDstr)

	if err != nil {
		utils.WriteError(w, http.StatusNotFound, "Request Failed", err.Error())
		return
	}

	utils.WriteSuccess(w, http.StatusOK, "size deleted successfully", deletedMaterial)

}

func (h *SizeHandler) GeySizeByID(w http.ResponseWriter, r *http.Request) {

	IDstr := chi.URLParam(r, "id")

	if r.Method != http.MethodGet {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Invalid Method", "Method Not Allowed")
		return
	}

	user, err := h.service.GetSizeByID(context.Background(), IDstr)

	if err != nil {
		utils.WriteError(w, http.StatusNotFound, "Request Failed", err.Error())
		return
	}

	utils.WriteSuccess(w, http.StatusOK, "size fetched successfully", user)

}
