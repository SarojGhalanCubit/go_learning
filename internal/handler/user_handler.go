package handler

import (
	"encoding/json"
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

	userID := r.Context().Value(middleware.UserIDKey)
	log.Println("USER ID :::: ",userID)

	users, err := h.service.GetUsers()
	if err != nil {
		utils.WriteError(w,http.StatusInternalServerError,err.Error(),"Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.WriteSuccess(w, http.StatusOK,"User fetched successfully", users)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteError(w,http.StatusMethodNotAllowed,"Invalid Method","Method Not Allowed")
		return
	}

	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {

		utils.WriteError(w,http.StatusInternalServerError,"Create User failed","Invalid Request Body")
		return
	}

	if userValidationErr := utils.ValidateUser(user.Name, user.Age, user.Phone,user.Email, user.Password); len(userValidationErr) > 0 {
		utils.WriteError(w,http.StatusUnprocessableEntity,"Validation Error",userValidationErr)
		return
	}
	
	hashedPassword,err := utils.HashPassword(user.Password); if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Password hashing failed",err.Error())
		return
	}

	user.Password = hashedPassword

 	created, err := h.service.CreateUser(user)
    	if err != nil {
		if err.Error() == "email already exists" || err.Error() == "phone already exists" {
			utils.WriteError(w,http.StatusConflict,err.Error(),"User creation failed",)
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
	utils.WriteSuccess(w, http.StatusCreated,"User created successfully", created)
}

func (h* UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost{
		utils.WriteError(w, http.StatusMethodNotAllowed,"Invalid Method","Method not allowed")
		return
	}

	var input struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.WriteError(w,http.StatusBadRequest,"Login Failed","Invalid Credentials" )
		return
	}

	if loginValidationErr := utils.ValidateLoginUser(input.Email,input.Password); len(loginValidationErr) > 0 {
		utils.WriteError(w,http.StatusUnprocessableEntity, "Validation Error",loginValidationErr)
		return
	}


	user,err := h.service.Login(input.Email,input.Password)
	if err != nil {
		utils.WriteError(w,http.StatusUnauthorized,"Login Failed","Invalid Credentials")
		return
	}
	token, err := utils.GenerateToken(user.ID)
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



	IDstr := chi.URLParam(r, "id")
	id,err := strconv.Atoi(IDstr)
	if err != nil {
		utils.WriteError(w,http.StatusNotFound, "Request Failed","Invalid user ID",)
		return 
	}
	

	log.Println("USER ID :::: ",userID)
	log.Println("USER ID :::: ",id)
	if userID != id {
		utils.WriteError(w,http.StatusNotFound, "Reqeust Failed", "No Permission for different user")
		return
	}

	user,err := h.service.GetUserByID(id)	
	if err != nil {
		utils.WriteError(w,http.StatusNotFound,"Request Failed",err.Error())
		return
	}

	utils.WriteSuccess(w, http.StatusOK, "User fetched successfully", user)

}


