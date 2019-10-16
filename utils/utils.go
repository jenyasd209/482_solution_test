package utils

import (
	"encoding/json"
	"log"
	"net/http"

	models "482.solution_test_task/models"
)

//CheckAuth - check basic auth
func CheckAuth(f func(w http.ResponseWriter, r *http.Request, user *models.User)) (h http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		username, password, authOK := r.BasicAuth()
		if authOK == false {
			http.Error(w, "Not authorized", 401)
			return
		}

		user, err := models.GetUserByUsername(username)
		if err != nil {
			http.Error(w, "Not found", 404)
			log.Println(err)
			return
		}

		if username != user.Username || password != user.Password {
			http.Error(w, "Not authorized", 401)
			return
		}

		f(w, r, &user)
	}
}

//Respond geting data
func Respond(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	w.WriteHeader(status)
}
