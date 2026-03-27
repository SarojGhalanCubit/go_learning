package handler

import (
	"encoding/json"
	"go-minimal/internal/config"
	"go-minimal/internal/middleware"
	"go-minimal/internal/model"
	"go-minimal/internal/service"
	"go-minimal/internal/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {

	if service == nil {
		panic("Service cannot be nil.")
	}

	return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := h.service.GetUsers()
	if err != nil {
		log.Println("err : ", err)
		utils.WriteError(w, http.StatusInternalServerError, err.Error(), "Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.WriteSuccess(w, http.StatusOK, "User fetched successfully", users)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Invalid Method", "Method Not Allowed")
		return
	}

	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {

		utils.WriteError(w, http.StatusInternalServerError, "Create User failed", "Invalid Request Body")
		return
	}

	if userValidationErr := utils.ValidateUser(user.Name, user.Age, user.Phone, user.Email, user.Password); len(userValidationErr) > 0 {
		utils.WriteError(w, http.StatusUnprocessableEntity, "Validation Error", userValidationErr)
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Password hashing failed", err.Error())
		return
	}

	user.Password = hashedPassword

	roleID := config.GetUserID()
	roleIdInt, err := strconv.Atoi(roleID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error(), "User creation failed")
		return
	}
	user.RoleID = roleIdInt

	created, err := h.service.CreateUser(user)
	if err != nil {
		if err.Error() == "email already exists" || err.Error() == "phone already exists" {
			utils.WriteError(w, http.StatusConflict, err.Error(), "User creation failed")
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
	utils.WriteSuccess(w, http.StatusCreated, "User created successfully", created)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

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
	token, err := utils.GenerateToken(user.ID, user.RoleID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Token generation failed", err.Error())
		return
	}

	utils.WriteSuccess(w, http.StatusOK, "Login successful", map[string]string{
		"token": token,
	})

}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(int)
	roleId := r.Context().Value(middleware.RoleIDKey).(int)
	IDstr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(IDstr)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, "Request Failed", "Invalid user ID")
		return
	}

	adminIDStr := config.GetAdminID()
	adminIDInt, err := strconv.Atoi(adminIDStr)
	if roleId != adminIDInt {

		if userID != id {
			utils.WriteError(w, http.StatusNotFound, "Reqeust Failed", "No Permission for different user")
			return
		}
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, "Request Failed", err.Error())
		return
	}

	utils.WriteSuccess(w, http.StatusOK, "User fetched successfully", user)

}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Invalid Method", "Method Not Allowed")
		return
	}

	var user model.UserResponse

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Update User failed", "Invalid Request Body")
		return
	}

	roleId := r.Context().Value(middleware.RoleIDKey).(int)
	userID := r.Context().Value(middleware.UserIDKey).(int)

	IDstr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(IDstr)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, "Request Failed", "Invalid user ID")
		return
	}

	adminIDStr := config.GetAdminID()
	adminIDInt, err := strconv.Atoi(adminIDStr)
	if roleId != adminIDInt {

		if userID != id {
			utils.WriteError(w, http.StatusNotFound, "Reqeust Failed", "No Permission for different user")
			return
		}
	}

	if userValidationErr := utils.ValidateUpdateUser(user.Name, user.Age, user.Phone, user.Email); len(userValidationErr) > 0 {
		utils.WriteError(w, http.StatusUnprocessableEntity, "Validation Error", userValidationErr)
		return
	}

	updated, err := h.service.UpdateUser(id, user)
	if err != nil {
		if err.Error() == "email already exists" || err.Error() == "phone already exists" {
			utils.WriteError(w, http.StatusConflict, err.Error(), "User update failed")
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
	utils.WriteSuccess(w, http.StatusCreated, "User updated successfully", updated)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(middleware.UserIDKey).(int)
	roleId := r.Context().Value(middleware.RoleIDKey).(int)

	if r.Method != http.MethodDelete {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Invalid Method", "Method Not Allowed")
		return
	}

	IDstr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(IDstr)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, "Request Failed", "Invalid user ID")
		return
	}

	adminIDStr := config.GetAdminID()
	adminIDInt, err := strconv.Atoi(adminIDStr)
	if roleId != adminIDInt {

		if userID != id {
			utils.WriteError(w, http.StatusNotFound, "Reqeust Failed", "No Permission for different user")
			return
		}
	}

	log.Println("PRINTER IDD :: ", id)
	deletedUser, err := h.service.DeleteUser(id)

	if err != nil {
		utils.WriteError(w, http.StatusNotFound, "Request Failed", err.Error())
		return
	}

	utils.WriteSuccess(w, http.StatusOK, "User Deleted successfully", deletedUser)

}
