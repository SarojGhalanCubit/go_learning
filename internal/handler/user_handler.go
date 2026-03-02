package handler 
import (
	"encoding/json"
	"net/http"
	"go-minimal/internal/model"
	"go-minimal/internal/service"
	"go-minimal/internal/utils"
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
		utils.WriteError(w,http.StatusInternalServerError,err.Error(),"Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
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
        	utils.WriteError(w, http.StatusInternalServerError, "Failed to create user", err.Error())
        	return
    	}	


	utils.WriteSuccess(w, http.StatusCreated,"User created successfully", created)
}



