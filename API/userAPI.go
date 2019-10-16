package api

import (
	"encoding/json"
	"log"
	"net/http"

	models "482.solution_test_task/models"
	utils "482.solution_test_task/utils"
)

//CreateUser - create new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		user := models.User{}

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "", 500)
			log.Println(err)
			return
		}

		if err := user.Construct(); err != nil {
			http.Error(w, err.Error(), 500)
			log.Println(err)
			return
		}

		if err := user.Insert(); err != nil {
			http.Error(w, "", 500)
			log.Println(err)
			return
		}

		utils.Respond(w, user, http.StatusOK)
	} else {
		utils.Respond(w, "Only POST method", 400)
	}
	return
}

//GetUser - return user by username
func GetUser(w http.ResponseWriter, r *http.Request, user *models.User) {
	if r.Method == http.MethodGet {
		utils.Respond(w, user, http.StatusOK)
	} else {
		utils.Respond(w, "Only GET method", 400)
	}
	return
}

//ChangeUser - change user fields
func ChangeUser(w http.ResponseWriter, r *http.Request, user *models.User) {
	if r.Method == http.MethodPut {

		updateUser := models.User{}

		err := json.NewDecoder(r.Body).Decode(&updateUser)
		if err != nil {
			http.Error(w, "", 400)
			log.Println(err)
			return
		}

		if err := updateUser.Construct(); err != nil {
			http.Error(w, err.Error(), 500)
			log.Println(err)
			return
		}

		if err := updateUser.Update(user.Username); err != nil {
			http.Error(w, "", 500)
			log.Println(err)
			return
		}
		utils.Respond(w, updateUser, http.StatusOK)
	} else {
		utils.Respond(w, "Only PUT method", 400)
	}
	return
}

//DeleteUser - delete user from store
func DeleteUser(w http.ResponseWriter, r *http.Request, user *models.User) {
	if r.Method == http.MethodDelete {
		deletedUser, err := user.Delete()
		if err != nil {
			http.Error(w, "", 500)
			log.Println(err)
			return
		}
		utils.Respond(w, deletedUser, http.StatusOK)
	} else {
		utils.Respond(w, "Only DELETE method", 400)
	}
	return
}
