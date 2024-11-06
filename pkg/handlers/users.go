package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "myapp/pkg/models"
)

func CreateUser(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var user models.User
        json.NewDecoder(r.Body).Decode(&user)
        
        // Insert user into DB
        _, err := db.Exec("INSERT INTO users (id, username, password, role) VALUES (?, ?, ?, ?)",
            user.ID, user.Username, user.Password, user.Role)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusCreated)
    }
}

func GetUser(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        params := mux.Vars(r)
        id := params["id"]

        var user models.User
        err := db.QueryRow("SELECT id, username, role FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Role)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }

        json.NewEncoder(w).Encode(user)
    }
}

// Similar functions for UpdateUser and DeleteUser
