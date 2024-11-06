package middlewares

import (
    "context"
    "net/http"
    "your_project/services"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        user, err := services.ValidateToken(token)
        if err != nil {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }

        ctx := context.WithValue(r.Context(), "user", user)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
