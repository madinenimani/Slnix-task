package models

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type User struct {
    ID       string
    Username string
    Password string // Should be hashed
    Role     string
}

type Org struct {
    ID      string
    Name    string
    AdminID string
}

type Folder struct {
    ID     string
    OrgID  string
    Name   string
    ParentID string
}

type File struct {
    ID       string
    FolderID string
    OrgID    string
    Name     string
    Path     string
    Permissions string
}

func InitDB(dataSourceName string) *sql.DB {
    db, err := sql.Open("sqlite3", dataSourceName)
    if err != nil {
        panic(err)
    }
    createTables(db)
    return db
}

func createTables(db *sql.DB) {
    db.Exec(`CREATE TABLE IF NOT EXISTS users (id TEXT PRIMARY KEY, username TEXT, password TEXT, role TEXT)`)
    db.Exec(`CREATE TABLE IF NOT EXISTS orgs (id TEXT PRIMARY KEY, name TEXT, admin_id TEXT)`)
    db.Exec(`CREATE TABLE IF NOT EXISTS folders (id TEXT PRIMARY KEY, org_id TEXT, name TEXT, parent_id TEXT)`)
    db.Exec(`CREATE TABLE IF NOT EXISTS files (id TEXT PRIMARY KEY, folder_id TEXT, org_id TEXT, name TEXT, path TEXT, permissions TEXT)`)
}
