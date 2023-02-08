package models

type User struct {
	Id        uint32 `json:"id"`
	FirstName string `json:"name_first"`
	LastName  string `json:"name_last"`
}
