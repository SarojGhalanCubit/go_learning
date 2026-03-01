package model

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
	Phone string `json:"phone_number"`
	Password string `json:"password"` // hide from returned JSON
}
 
