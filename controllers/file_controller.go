package controllers

import (
    "net/http"
    "your_project/services"
    "github.com/gorilla/mux"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
    file, _, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "File upload failed", http.StatusBadRequest)
        return
    }
    defer file.Close()

    fileID, err := services.SaveFile(file)
    if err != nil {
        http.Error(w, "Could not save file", http.StatusInternalServerError)
        return
    }

    w.Write([]byte(fileID))
}

func GetFile(w http.ResponseWriter, r *http.Request) {
    fileID := mux.Vars(r)["id"]

    file, err := services.RetrieveFile(fileID)
    if err != nil {
        http.Error(w, "File not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Disposition", "attachment; filename="+file.Name)
    http.ServeFile(w, r, file.Path)
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {
    fileID := mux.Vars(r)["id"]

    err := services.DeleteFile(fileID)
    if err != nil {
        http.Error(w, "File deletion failed", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
