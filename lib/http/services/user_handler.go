package modularHTTP

import (
	. "ModularHTTPGo/types"
	. "ModularHTTPGo/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type NewUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUserHandler(db *sql.DB) Handler {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func PostUserHandler(db *sql.DB) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		SetContentJson(w)

		user := NewUserRequest{}
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		hashPassword, err := GenerateSecurePassword(user.Password)

		if err != nil {
			http.Error(w, http.StatusText(500), 500)

			return
		}

		user.Password = hashPassword
		fmt.Println(user.Password)

		value, err := db.Exec(fmt.Sprintf("INSERT INTO users (username, password) VALUES ($1, $2)"), user.Username, user.Password)
		if err != nil {

			http.Error(w, http.StatusText(500), 500)
			return
		}

		aff, err := value.RowsAffected()
		if aff < 1 {

			http.Error(w, http.StatusText(500), 500)
			return
		}

		if err != nil {

			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(BaseResponse{
			Status: 200,
			Alert:  map[string]interface{}{},
			Data:   fmt.Sprintf("username : %v created successfully", user.Username),
		})

	}
}
func UpdateUserHandler(db *sql.DB) Handler {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func DeleteUserHandler(db *sql.DB) Handler {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func GetUserByIdHandler(db *sql.DB) Handler {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
