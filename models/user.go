package models

import (
	"database/sql"
	"errors"
	"log"
	"unicode"

	dbConnection "482.solution_test_task/db"
)

//User struct
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	About    string `json:"about"`
}

//Construct user
func (user *User) Construct() (err error) {
	if err = user.CheckUsername(); err != nil {
		return
	}
	return
}

//CheckUsername - check username if it have lowwer case or number or "_"
func (user *User) CheckUsername() (err error) {
	if CheckExistUsername(user.Username) {
		err = errors.New("Username alreade exist")
		return
	}
	runeSlice := []rune(user.Username)
	for i := 0; i < len(runeSlice); i++ {
		if unicode.IsLower(runeSlice[i]) || unicode.IsNumber(runeSlice[i]) || runeSlice[i] == '_' {
			continue
		} else {
			err = errors.New("Invalid character")
			return
		}
	}
	return
}

//CheckExistUsername - return true if exist
func CheckExistUsername(username string) bool {
	statement := `SELECT username FROM user WHERE username = ?`
	err := dbConnection.Db.QueryRow(statement, username).Scan(&username)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return false
	}
	return true
}

//Insert new user in table user
func (user *User) Insert() (err error) {
	statement := `INSERT INTO USER (username, password, about) VALUES ($1, $2, $3);`
	stmt, err := dbConnection.Db.Prepare(statement)
	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(user.Username, user.Password, user.About)
	if err != nil {
		return
	}
	return
}

//Update user in table user
func (user *User) Update(username string) (err error) {
	statement := `UPDATE USER SET username = $1, password = $2, about = $3 WHERE username = $4`
	stmt, err := dbConnection.Db.Prepare(statement)
	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(user.Username, user.Password, user.About, username)
	if err != nil {
		return
	}
	return
}

//GetUserByUsername - return user(struct User) by username(string)
func GetUserByUsername(username string) (user User, err error) {
	err = dbConnection.Db.QueryRow(`SELECT id, username, password, about FROM user
		WHERE username = $1`, username).Scan(&user.ID, &user.Username, &user.Password, &user.About)
	if err != nil {
		return
	}
	return
}

//Delete user from storage
func (user *User) Delete() (deletedUser User, err error) {
	deletedUser, err = GetUserByUsername(user.Username)
	if err != nil {
		return
	}
	statement := "DELETE FROM user WHERE username = $1"
	stmt, err := dbConnection.Db.Prepare(statement)
	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(user.Username)
	return
}
