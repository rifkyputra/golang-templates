package modularHTTP

import (
	. "ModularHTTPGo/types"
	. "ModularHTTPGo/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags users
// @ID get-user-by-id
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 404 {string} string "User not found"
// @Router /users/{id} [get]
func PostAuthHandler(db *sql.DB) Handler {
	return func(w http.ResponseWriter, r *http.Request) {

		SetContentJson(w)
		user := User{}

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		rows, err := db.Query("SELECT username, password FROM users WHERE username =$1", user.Username)

		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			fmt.Println(err)

			return
		}

		defer rows.Close()

		userQuery := User{}
		has := rows.Next()
		if !has {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			fmt.Println(has)
			fmt.Println(userQuery)

			return
		}
		err = rows.Scan(&userQuery.Username, &userQuery.Password)

		if err != nil {

			http.Error(w, http.StatusText(400), 400)
			fmt.Println(err)
			fmt.Println(userQuery)

			return
		}

		identical := IsPasswordIdentical(userQuery.Password, user.Password) // Password is incorrect, handle error

		if !identical {
			fmt.Println(err)
			http.Error(w, http.StatusText(400), 400)
			return
		}

		jwt, err := GenerateJWT([]byte("my_secret_key"))

		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(BaseResponse{
			Status: 200,
			Alert:  map[string]interface{}{},
			Data:   jwt,
		})

	}
}
