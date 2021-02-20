package shared

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/vasuvanka/todo-app/backend/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

// CreateJWT - creates jwt token
func CreateJWT(id, name, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = name
	claims["role"] = role
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	return token.SignedString([]byte(JWTSecret))
}

//GetTokenFromHeader - get jwt token from header
func GetTokenFromHeader(r *http.Request) (string, error) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return "", errors.New("Unauthorized access")
	}

	authTokens := strings.Split(auth, " ")
	if len(authTokens) != 2 {
		return "", errors.New("Unauthorized access")
	}

	return authTokens[1], nil
}

//ValidateJWT - validate and get jwt payload
func ValidateJWT(jwToken string) (jwt.Claims, error) {
	token, err := jwt.Parse(jwToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", errors.New("invalid token")
		}
		return []byte(JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	if err = token.Claims.Valid(); err != nil {
		return nil, err
	}

	return token.Claims, nil
}

//HashPassword - hash password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

// CheckPasswordHash - validate password aganist hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// SendJSON - send json payload
func SendJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		SendError(w, models.Response{Status: http.StatusBadRequest, Message: "invalid request"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// SendError - send json error
func SendError(w http.ResponseWriter, m models.Response) {
	response, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(m.Status)
	w.Write(response)
}
