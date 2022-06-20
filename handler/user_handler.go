package user_handler

import (
	"CRUDRestApiDatabase/database"
	user "CRUDRestApiDatabase/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type UserHandlerInterface interface {
	UserHandler(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	db *sql.DB
}

func NewUserHandler(db *sql.DB) UserHandlerInterface {
	return &UserHandler{db: db}
}

func (h *UserHandler) UserHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]
	switch r.Method {
	case http.MethodGet:
		if id != "" {
			h.getUserByIdHandler(w, r, id)
		} else {
			h.getUserHandlerAll(w, r)
		}
	case http.MethodPost:
		h.createUserHandler(w, r)
	case http.MethodPut:
		h.updateUserHandler(w, r, id)
	case http.MethodDelete:
		h.deleteUserHandler(w, r, id)
	}
}

func (h *UserHandler) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var userss = user.User{}
	json.NewDecoder(r.Body).Decode(&userss)
	userss.Created_at = time.Now()
	userss.Updated_at = time.Now()
	sqlInject :=
		`INSERT INTO users(
		username, email, pass, age_user, division, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7)
		Returning id ;`

	err := database.Db.QueryRow(sqlInject,
		userss.Username,
		userss.Email,
		userss.Password,
		userss.Age_user,
		userss.Division,
		userss.Created_at,
		userss.Updated_at,
	).Scan(&userss.Id)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(fmt.Sprint("User Created")))
	return
}
func (h *UserHandler) getUserByIdHandler(w http.ResponseWriter, r *http.Request, id string) {
	var result = []user.User{}
	if id != "" {
		sqlGet := `Select * from users where id = $1`
		rows, err := database.Db.Query(sqlGet, id)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			var userss = user.User{}
			if err = rows.Scan(&userss.Id,
				&userss.Username,
				&userss.Email,
				&userss.Password,
				&userss.Age_user,
				&userss.Division,
				&userss.Created_at,
				&userss.Updated_at,
			); err != nil {
				fmt.Println("No Data", err)
			}
			result = append(result, userss)
		}
		jsonData, _ := json.Marshal(&result)
		w.Header().Add("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func (h *UserHandler) getUserHandlerAll(w http.ResponseWriter, r *http.Request) {
	var result = []user.User{}
	sqlGet := "Select * from users;"
	rows, err := database.Db.Query(sqlGet)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var userss = user.User{}
		if err = rows.Scan(
			&userss.Id,
			&userss.Username,
			&userss.Email,
			&userss.Password,
			&userss.Age_user,
			&userss.Division,
			&userss.Created_at,
			&userss.Updated_at,
		); err != nil {
			fmt.Println("No Data", err)
		}
		result = append(result, userss)
	}
	jsonData, _ := json.Marshal(&result)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}

func (h *UserHandler) updateUserHandler(w http.ResponseWriter, r *http.Request, id string) {
	if id != "" {
		var userss = user.User{}
		json.NewDecoder(r.Body).Decode(&userss)
		userss.Created_at = time.Now()
		userss.Updated_at = time.Now()
		sqlUpdate := `UPDATE users SET
		username = $2,
		email = $3,
		pass= $4,
		age_user = $5,
		division= $6,
		created_at = $7,
		updated_at = $8
		WHERE id = $1`
		res, err := database.Db.Exec(sqlUpdate, id,
			userss.Username,
			userss.Email,
			userss.Password,
			userss.Age_user,
			userss.Division,
			userss.Created_at,
			userss.Updated_at,
		)
		if err != nil {
			panic(err)
		}
		count, err := res.RowsAffected()
		if err != nil {
			panic(err)
		}
		w.Write([]byte(fmt.Sprint("Updated Data", count)))
		return
	}

}

func (h *UserHandler) deleteUserHandler(w http.ResponseWriter, r *http.Request, id string) {
	sqlDelete := `DELETE from users WHERE id = $1`
	if index, err := strconv.Atoi(id); err == nil {
		res, err := database.Db.Exec(sqlDelete, index)
		if err != nil {
			panic(err)
		}
		count, err := res.RowsAffected()
		if err != nil {
			panic(err)
		}

		w.Write([]byte(fmt.Sprint("Deleted Data", count)))
		return
	}
}
