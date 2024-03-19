package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/puvadon-artmit/gofiber-template/config"
	"github.com/puvadon-artmit/gofiber-template/model"
	"golang.org/x/crypto/bcrypt"
)

func Encode(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes)
}

func Compare(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJwt(user model.User) (string, error) {
	var privateKey = config.GetEnvConfig("SECRET_KEY")
	claims := jwt.MapClaims{
		"id":  user.UserID,
		"exp": time.Now().Add(time.Hour * 10).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(privateKey))
	if err != nil {
		return "Fail", err
	}
	return t, nil
}
