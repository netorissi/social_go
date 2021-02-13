package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func Hash(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}

func ValidHash(hash, pass string) (err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass)); err != nil {
		err = errors.New("this login is unauthorized")
	}
	return err
}

func CreateToken(userID uint64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.SecretKey)
}

func ValidToken(r *http.Request) error {
	tokenStr := getAuthToken(r)
	token, err := jwt.Parse(tokenStr, validKeyToken)
	if err != nil {
		return errors.New("token is invalid")
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token is invalid")
}

func getAuthToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	tokenSplit := strings.Split(token, " ")

	if len(tokenSplit) == 2 {
		return tokenSplit[1]
	}

	return ""
}

func getAuthUserID(r *http.Request) (uint64, error) {
	tokenStr := getAuthToken(r)
	token, err := jwt.Parse(tokenStr, validKeyToken)
	if err != nil {
		return 0, errors.New("token is invalid")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 64)
		if err != nil {
			return 0, errors.New("token is invalid")
		}
		return userID, nil
	}

	return 0, errors.New("token is invalid")
}

func validKeyToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Method signing token is incorrect. %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
