package model 

type RoleResponse struct {
	ID int `json:"id"`
	RoleName string `json:"role_name"`
	IsActive bool `json:"is_active"`
}
