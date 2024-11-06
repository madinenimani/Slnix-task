package storage

import (
    "io"
    "os"
    "path/filepath"
)

func SaveLocalFile(file io.Reader, filename string) (string, error) {
    path := filepath.Join("storage", filename)
    outFile, err := os.Create(path)
    if err != nil {
        return "", err
    }
    defer outFile.Close()

    _, err = io.Copy(outFile, file)
    if err != nil {
        return "", err
    }
    return path, nil
}
