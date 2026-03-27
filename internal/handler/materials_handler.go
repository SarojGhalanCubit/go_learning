package handler

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go-minimal/internal/model"
	"go-minimal/internal/service"
	"go-minimal/internal/utils"
	materialValidate "go-minimal/internal/utils/validate"
	"log"
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

func (h *MaterialHandler) CreateMaterial(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Invalid method", "Method not allowed")
		return
	}

	var material model.CreateMaterial

	err := json.NewDecoder(r.Body).Decode(&material)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "create material failed", "Invalid request body")
	}

	if userValidationErr := materialValidate.ValidateMaterial(material.Name); len(userValidationErr) > 0 {
		utils.WriteError(w, http.StatusUnprocessableEntity, "Validation Error", userValidationErr)
		return
	}

	createdMaterial, err := h.service.CreateMaterial(context.Background(), material)
	log.Println("server error", err)

	if err != nil {
		if err.Error() == "material name already exists" {
			utils.WriteError(w, http.StatusConflict, err.Error(), "material creation failed")
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

	utils.WriteSuccess(w, http.StatusOK, "material created successfully", createdMaterial)

}

func (h *MaterialHandler) UpdateMaterial(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Invalid Method", "Method Not Allowed")
		return
	}

	var material model.CreateMaterial

	if err := json.NewDecoder(r.Body).Decode(&material); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Update material failed", "Invalid Request Body")
		return
	}

	if userValidationErr := materialValidate.ValidateMaterial(material.Name); len(userValidationErr) > 0 {
		utils.WriteError(w, http.StatusUnprocessableEntity, "Validation Error", userValidationErr)
		return
	}

	IDstr := chi.URLParam(r, "id")

	updated, err := h.service.UpdateMaterial(context.Background(), IDstr, material)
	log.Println("ERR :: ", err)
	if err != nil {
		if err.Error() == "material name already exists" || err.Error() == "requested material did not exist" {
			utils.WriteError(w, http.StatusConflict, err.Error(), "material update failed")
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
	utils.WriteSuccess(w, http.StatusCreated, "Material updated successfully", updated)
}

func (h *MaterialHandler) DeleteMaterial(w http.ResponseWriter, r *http.Request) {

	IDstr := chi.URLParam(r, "id")
	if r.Method != http.MethodDelete {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Invalid Method", "Method Not Allowed")
		return
	}

	deletedMaterial, err := h.service.DeleteMaterialById(context.Background(), IDstr)

	if err != nil {
		utils.WriteError(w, http.StatusNotFound, "Request Failed", err.Error())
		return
	}

	utils.WriteSuccess(w, http.StatusOK, "material deleted successfully", deletedMaterial)

}

func (h *MaterialHandler) GeyByMaterialID(w http.ResponseWriter, r *http.Request) {

	IDstr := chi.URLParam(r, "id")

	if r.Method != http.MethodGet {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Invalid Method", "Method Not Allowed")
		return
	}

	user, err := h.service.GeyByMaterialID(context.Background(), IDstr)

	if err != nil {
		utils.WriteError(w, http.StatusNotFound, "Request Failed", err.Error())
		return
	}

	utils.WriteSuccess(w, http.StatusOK, "Material fetched successfully", user)

}
