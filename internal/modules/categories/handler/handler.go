package categoriesHandler

import (
	"context"
	categoriesService "go-minimal/internal/modules/categories/service"
	"go-minimal/internal/utils"
	"net/http"
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
	}

	sizes, err := s.service.GetAllCategories(context.Background())

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error(), "Internal Server Error")
	}

	w.Header().Set("Content-Type", "application/json")
	utils.WriteSuccess(w, http.StatusOK, "categories fetched successfully", sizes)

}
