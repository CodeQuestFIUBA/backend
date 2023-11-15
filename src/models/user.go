package models

type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Active  bool   `json:"active,omitempty"`
}
