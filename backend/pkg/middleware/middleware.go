package middleware

import (
    "net/http"
    "strings"
    "backend/internal/auth"
)

func JWTMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") 
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")


        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }


        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte("Autorização necessária"))
            return
        }


        tokenParts := strings.Split(authHeader, " ")
        if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte("Formato de token inválido"))
            return
        }

        tokenString := tokenParts[1]
        claims, err := auth.ValidateToken(tokenString)
        if err != nil {
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte("Token inválido"))
            return
        }

        r.Header.Set("username", claims.Username)

        next.ServeHTTP(w, r)
    })
}