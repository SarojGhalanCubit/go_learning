package productHandler

import (
	"context"
	"encoding/json"
	productModel "go-minimal/internal/modules/products/model"
	productService "go-minimal/internal/modules/products/service"
	"go-minimal/internal/utils"

	productValidate "go-minimal/internal/utils/validate"
	"net/http"
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
	var req productModel.CreateProduct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request", "JSON decode failed")
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
		utils.WriteError(w, http.StatusInternalServerError, "Failed to create product", err.Error())
		return
	}

	utils.WriteSuccess(w, http.StatusCreated, "Product created successfully", product)
}
