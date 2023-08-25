package services

import (
	"dudu/models"
	"encoding/json"
	"fmt"
	"log"
	// jwt "github.com/dgrijalva/jwt-go"
)

// 定义userService
type UserService struct {
}

func (*UserService) CreateNewUser(body []byte) int16 {
	user := models.User{}
	json.Unmarshal(body, &user)

	fmt.Println(body)
	has := user.VerifyUser(user)
	if !has {
		return 0
	}

	id, err := user.Insert(user)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return id
}

// 签名密钥
const sign_key = "hello jwt"

func (*UserService) UserLogin(email, password string) bool {

	// token, _ := (email, password)

	// Println(token)

	return true
}

func init() {

}
