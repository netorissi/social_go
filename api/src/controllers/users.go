package controllers

import "net/http"

// GetUser -
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User."))
}

// CreateUser -
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create User."))
}

// UpdateUser -
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User."))
}

// DeleteUser -
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User."))
}