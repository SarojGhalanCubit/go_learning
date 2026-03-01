package handler 
import (
	"encoding/json"
	"net/http"
	"log"
	"go-minimal/internal/model"
	"go-minimal/internal/service"
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}


func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}


	log.Println("user info :: ",user)
	created, err := h.service.CreateUser(user)

	// if err != nil {
	// log.Println("CreateUser Error:", err) // Backend debugging
	// 	http.Error(w, "Database error", http.StatusInternalServerError)
	// 	return
	// }
	//
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)

}
