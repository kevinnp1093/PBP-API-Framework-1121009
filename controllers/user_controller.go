package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"

	"github.com/your-username/your-app/db"
	"github.com/your-username/your-app/model"
)

func GetUser(params martini.Params, w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	err := db.Conn.QueryRow("SELECT id, username, email FROM users WHERE id=?", params["id"]).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	email := r.FormValue("email")

	result, err := db.Conn.Exec("INSERT INTO users (username, email) VALUES (?, ?)", username, email)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	user := model.User{
		ID:       int(id),
		Username: username,
		Email:    email,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(params martini.Params, w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(params["id"])
	username := r.FormValue("username")
	email := r.FormValue("email")

	result, err := db.Conn.Exec("UPDATE users SET username=?, email=? WHERE id=?", username, email, id)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	rowsAffected, _
}