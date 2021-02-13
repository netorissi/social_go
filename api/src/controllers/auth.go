package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	db, err := database.Init()
	if err != nil {
		utils.AppError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewRepositoryUsers(db)
	userExist, err := userRepository.FindByEmail(user.Email)
	if err != nil {
		utils.AppError(w, http.StatusBadRequest, err)
		return
	}

	if err := auth.ValidHash(userExist.Password, user.Password); err != nil {
		utils.AppError(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.CreateToken(userExist.ID)
	if err != nil {
		utils.AppError(w, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte(token))
}
