package handler 
import (
	"encoding/json"
	"net/http"
	"go-minimal/model"
	"go-minimal/service"
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
	var users []model.User =  h.service.GetUsers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}


func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Method not allowed",
		})

		return
	}

	var user model.User

	// Decode json
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid Request Bodyk j",
		})

		return
	}

	if user.Name == "" {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Name is required",
		})

		return 
	}

	if user.Age < 18 {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":"You are under age",
		})
		return 
	}


	created := h.service.CreateUser(user)

	// Response
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)


	json.NewEncoder(w).Encode(created)

}
