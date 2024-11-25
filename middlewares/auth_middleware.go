package middlewares

import (
	"BankTellerAPI/database"
	"BankTellerAPI/utils"
	"context"
	"net/http"
	"strings"
)

type key int

const userIDKey key = 0

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "Malformed token", http.StatusUnauthorized)
			return
		}

		isBlacklisted, err := database.IsTokenBlacklisted(tokenString)
		if err != nil {
			http.Error(w, "Error checking token blacklist", http.StatusInternalServerError)
			return
		}
		if isBlacklisted {
			http.Error(w, "Token is blacklisted", http.StatusUnauthorized)
			return
		}

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		userID, ok := claims["userID"].(string)
		if !ok {
			http.Error(w, "UserID not found in token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userIDKey, userID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}
}

// GetUserID extracts userID from context
func GetUserID(r *http.Request) string {
	userID, _ := r.Context().Value(userIDKey).(string)
	return userID
}
