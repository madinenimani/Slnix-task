package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "myapp/pkg/auth"
    "myapp/pkg/handlers"
    "myapp/pkg/models"
)

func main() {
    err := godotenv.Load("config.yaml")
    if err != nil {
        log.Fatalf("Error loading config file: %v", err)
    }

    db := models.InitDB("db.sqlite")

    router := mux.NewRouter()

    // User management endpoints
    router.HandleFunc("/users", handlers.CreateUser(db)).Methods("POST")
    router.HandleFunc("/users/{id}", handlers.GetUser(db)).Methods("GET")
    router.HandleFunc("/users/{id}", handlers.UpdateUser(db)).Methods("PUT")
    router.HandleFunc("/users/{id}", handlers.DeleteUser(db)).Methods("DELETE")

    // Organization and folder management
    router.HandleFunc("/orgs", handlers.CreateOrg(db)).Methods("POST")
    router.HandleFunc("/orgs/{orgID}/folders", handlers.CreateFolder(db)).Methods("POST")
    router.HandleFunc("/orgs/{orgID}/folders/{folderID}", handlers.DeleteFolder(db)).Methods("DELETE")

    // File management
    router.HandleFunc("/orgs/{orgID}/folders/{folderID}/files", handlers.UploadFile(db)).Methods("POST")
    router.HandleFunc("/orgs/{orgID}/folders/{folderID}/files/{fileID}", handlers.DownloadFile(db)).Methods("GET")
    router.HandleFunc("/orgs/{orgID}/folders/{folderID}/files/{fileID}", handlers.DeleteFile(db)).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8080", router))
}
