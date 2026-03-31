package categoriesHandler

import (
	"context"
	"encoding/json"
	categoryModel "go-minimal/internal/modules/categories/model"
	categoriesService "go-minimal/internal/modules/categories/service"
	"go-minimal/internal/utils"
	ValidateCategory "go-minimal/internal/utils/validate"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type CategoriesHandler struct {
	service *categoriesService.CategoriesService
}

func NewCategoriesHandler(s *categoriesService.CategoriesService) *CategoriesHandler {
	return &CategoriesHandler{
		service: s,
	}
}

func (s *CategoriesHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Invalid request method ", "method not allowed")
		return
	}

	sizes, err := s.service.GetAllCategories(context.Background())

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error(), "Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.WriteSuccess(w, http.StatusOK, "categories fetched successfully", sizes)

}

func (h *CategoriesHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Invalid method", "Method not allowed")
		return
	}

	var category categoryModel.CreateCategory

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "create category failed", "Invalid request body")
		return
	}

	if userValidationErr := ValidateCategory.ValidateMaterial(category.Name); len(userValidationErr) > 0 {

		log.Println("validate err :", userValidationErr)
		utils.WriteError(w, http.StatusUnprocessableEntity, "Validation Error", userValidationErr)
		return
	}

	createdCategory, err := h.service.CreateCategory(context.Background(), category)
	if err != nil {
		if err.Error() == "category name already exists" || err.Error() == "this slug is already taken" {
			utils.WriteError(w, http.StatusConflict, err.Error(), "category creation failed")
			return
		}
		utils.WriteError(
			w,
			http.StatusInternalServerError,
			"Internal server error",
			"category creation failed",
		)

		return
	}

	utils.WriteSuccess(w, http.StatusOK, "category created successfully", createdCategory)

}

func (h *CategoriesHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Invalid Method", "Method Not Allowed")
		return
	}

	var category categoryModel.CreateCategory

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Update category failed", "Invalid Request Body")
		return
	}

	if userValidationErr := ValidateCategory.ValidateMaterial(category.Name); len(userValidationErr) > 0 {
		utils.WriteError(w, http.StatusUnprocessableEntity, "Validation Error", userValidationErr)
		return
	}

	IDstr := chi.URLParam(r, "id")

	updated, err := h.service.UpdateCategory(context.Background(), IDstr, category)
	if err != nil {
		if err.Error() == "category name already exists" || err.Error() == "requested category did not exist" || err.Error() == "this slug is already taken" {
			utils.WriteError(w, http.StatusConflict, err.Error(), "material update failed")
			return
		}

		utils.WriteError(
			w,
			http.StatusInternalServerError,
			"Internal server error",
			err.Error(),
		)

		return
	}
	utils.WriteSuccess(w, http.StatusCreated, "category updated successfully", updated)
}
func (h *CategoriesHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {

	IDstr := chi.URLParam(r, "id")
	if r.Method != http.MethodDelete {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Invalid Method", "Method Not Allowed")
		return
	}

	deletedMaterial, err := h.service.DeleteCategoryById(context.Background(), IDstr)

	if err != nil {
		utils.WriteError(w, http.StatusNotFound, "Request Failed", err.Error())
		return
	}

	utils.WriteSuccess(w, http.StatusOK, "category deleted successfully", deletedMaterial)

}

func (h *CategoriesHandler) GeyByCategoryID(w http.ResponseWriter, r *http.Request) {

	IDstr := chi.URLParam(r, "id")

	if r.Method != http.MethodGet {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Invalid Method", "Method Not Allowed")
		return
	}

	user, err := h.service.GeyByCategoryID(context.Background(), IDstr)

	if err != nil {
		utils.WriteError(w, http.StatusNotFound, "Request Failed", err.Error())
		return
	}

	utils.WriteSuccess(w, http.StatusOK, "category fetched successfully", user)

}
