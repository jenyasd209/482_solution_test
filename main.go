package main

import (
	"net/http"

	api "482.solution_test_task/API"
	utils "482.solution_test_task/utils"
)

func main() {
	http.HandleFunc("/create_user", api.CreateUser)
	http.HandleFunc("/get_user", utils.CheckAuth(api.GetUser))
	http.HandleFunc("/change_user", utils.CheckAuth(api.ChangeUser))
	http.HandleFunc("/delete_user", utils.CheckAuth(api.DeleteUser))
	http.ListenAndServe(":8080", nil)
}
