package authHandler

import (
	"encoding/json"
	authLoginService "go-minimal/internal/modules/auth/service"
	"go-minimal/internal/utils"
	"log"
	"net/http"
)

type AuthHandler struct {
	service *authLoginService.AuthService
}

func NewAuthHandler(service *authLoginService.AuthService) *AuthHandler {

	if service == nil {
		panic("Service cannot be nil.")
	}

	return &AuthHandler{service: service}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Invalid Method", "Method not allowed")
		return
	}

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Login Failed", "Invalid Credentials")
		return
	}

	if loginValidationErr := utils.ValidateLoginUser(input.Email, input.Password); len(loginValidationErr) > 0 {
		utils.WriteError(w, http.StatusUnprocessableEntity, "Validation Error", loginValidationErr)
		return
	}

	user, err := h.service.Login(input.Email, input.Password)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "Login Failed", "Invalid Credentials")
		return
	}

	log.Println("USER DATA ::: ", input.Email, input.Password, user.RoleID)
	token, err := utils.GenerateToken(user.ID, user.RoleID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Token generation failed", err.Error())
		return
	}

	utils.WriteSuccess(w, http.StatusOK, "Login successful", map[string]string{
		"token": token,
	})

}
