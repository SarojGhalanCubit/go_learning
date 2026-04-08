package productHandler

import (
	"context"
	productModel "go-minimal/internal/modules/products/model"
	productService "go-minimal/internal/modules/products/service"
	"go-minimal/internal/utils"
	productValidate "go-minimal/internal/utils/validate"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	svc *productService.ProductService
}

func NewProductHandler(svc *productService.ProductService) *ProductHandler {
	return &ProductHandler{svc: svc}
}

func (handler *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteError(w, http.StatusMethodNotAllowed, "invalid method", "method not allowed")
		return
	}
	products, err := handler.svc.GetAllProducts(context.Background())

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "products fetch failed", "internal server error")
		return
	}

	utils.WriteSuccess(w, http.StatusOK, "products fetched successfully", products)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteError(w, http.StatusMethodNotAllowed, "invalid method", "method not allowed")
		return
	}
	var req productModel.CreateProduct
	// Use our new friendly decoder
	if err := utils.DecodeJSON(w, r, &req); err != nil {
		// Error response is already sent by the utility, just stop here
		return
	}

	// Use the validator we wrote
	valErrors := productValidate.ValidateProduct(req.Name, req.Description, req.Quantity, req.CategoryID, req.MaterialID)
	if len(valErrors) > 0 {
		utils.WriteError(w, http.StatusUnprocessableEntity, "Validation Error", valErrors)
		return
	}

	product, err := h.svc.CreateProduct(r.Context(), req)
	if err != nil {
		if err.Error() == "product name already exists" {
			utils.WriteError(w, http.StatusConflict, err.Error(), "product creation failed")
			return
		}

		utils.WriteError(w, http.StatusInternalServerError, "Failed to create product", err.Error())
		return
	}

	utils.WriteSuccess(w, http.StatusCreated, "Product created successfully", product)
}

func (h *ProductHandler) UpdateProductByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		utils.WriteError(w, http.StatusMethodNotAllowed, "invalid method", "method not allowed")
		return
	}
	var req productModel.CreateProduct
	// Use our new friendly decoder
	if err := utils.DecodeJSON(w, r, &req); err != nil {
		// Error response is already sent by the utility, just stop here
		return
	}
	valErrors := productValidate.ValidateProduct(req.Name, req.Description, req.Quantity, req.CategoryID, req.MaterialID)
	if len(valErrors) > 0 {
		utils.WriteError(w, http.StatusUnprocessableEntity, "Validation Error", valErrors)
		return
	}

	IDstr := chi.URLParam(r, "id")

	product, err := h.svc.UpdateProductByID(r.Context(), req, IDstr)
	if err != nil {
		if err.Error() == "product name already exists" {
			utils.WriteError(w, http.StatusConflict, err.Error(), "product creation failed")
			return
		}

		utils.WriteError(w, http.StatusInternalServerError, "Failed to update product", err.Error())
		return
	}

	utils.WriteSuccess(w, http.StatusCreated, "Product updated successfully", product)

}
