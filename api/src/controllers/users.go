package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// GetUsers -
func GetUsers(w http.ResponseWriter, r *http.Request) {
	param := strings.ToLower(r.URL.Query().Get("search"))

	db, err := database.Init()
	if err != nil {
		utils.AppError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewRepositoryUsers(db)

	users, err := userRepository.Find(param)
	if err != nil {
		utils.AppError(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, users)
}

// GetUserByID -
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		utils.AppError(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Init()
	if err != nil {
		utils.AppError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewRepositoryUsers(db)

	user, err := userRepository.FindByID(userID)
	if err != nil {
		utils.AppError(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusOK, user)
}

// CreateUser -
func CreateUser(w http.ResponseWriter, r *http.Request) {
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.AppError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(req, &user); err != nil {
		utils.AppError(w, http.StatusBadRequest, err)
		return
	}

	if err = user.BeforeCreate(); err != nil {
		utils.AppError(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Init()
	if err != nil {
		utils.AppError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewRepositoryUsers(db)

	user.ID, err = userRepository.Create(user)
	if err != nil {
		utils.AppError(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusCreated, user)
}

// UpdateUser -
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User."))
}

// DeleteUser -
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User."))
}
