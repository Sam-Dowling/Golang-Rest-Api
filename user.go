package main

//User holds info for single record in user table
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}
