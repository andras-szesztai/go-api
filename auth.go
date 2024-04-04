package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func WithJWTAuth(next http.HandlerFunc, store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := validateToken(getTokenFromRequest(r))
		if err != nil {
			WriteJson(w, http.StatusUnauthorized, ErrorResponse{Error: "Invalid token"})
			return
		}
		if !token.Valid {
			WriteJson(w, http.StatusUnauthorized, ErrorResponse{Error: "Invalid token"})
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		userId := claims["userId"].(string)

		_, err = store.GetUserById(userId)
		if err != nil {
			WriteJson(w, http.StatusUnauthorized, ErrorResponse{Error: "Invalid token"})
			return
		}

		next(w, r)
	}
}

func getTokenFromRequest(r *http.Request) string {
	if tokenAuth := r.Header.Get("Authorization"); tokenAuth != "" {
		return tokenAuth
	}
	return ""
}

func validateToken(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CreateAndSetAuthCookie(userID int64, w http.ResponseWriter) (string, error) {
	secret := []byte(Env.JWTSecret)
	token, err := CreateJWT(secret, userID)
	if err != nil {
		return "", err
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "Authorization",
		Value: token,
	})

	return token, nil
}

func CreateJWT(secret []byte, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(int(userID)),
		"expiresAt": time.Now().Add(time.Hour * 24 * 120).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, err
}
