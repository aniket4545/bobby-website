package admin

import (
	"errors"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

var err = errors.New("Invalid input. Please enter valid email and password")

type admin struct {
	name        string
	email       string
	password    string
	accessToken string
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func hashAndSalt(password []byte) string {
	encrypted, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic("While creating password: " + err.Error())
	}
	return string(encrypted)
}

//isValidPassword will validate password for user
func (u *admin) isValidPassword(password []byte) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.password), password); err != nil {
		return false
	}
	return true
}

//GenerateAccessToken will generate access token for user
func (u *admin) generateAccessToken() {
	u.accessToken = getAccessToken(30)
}

func getAccessToken(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//ValidateRequest will validate request and will return token
func ValidateRequest(email string, password string) (*string, error) {
	if email != ADMIN.email || !ADMIN.isValidPassword([]byte(password)) {
		return nil, err
	}
	ADMIN.generateAccessToken()
	token := ADMIN.accessToken
	return &token, nil
}

//DestroyAccessToken will be called on signout call
func (u *admin) DestroyAccessToken() {
	u.accessToken = ""
}

//RefreshToken will regenerate the token and will send it to client
func RefreshToken() *string {
	ADMIN.generateAccessToken()
	return &ADMIN.accessToken
}
