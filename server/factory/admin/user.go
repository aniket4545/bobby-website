package admin

import (
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

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

func (u *admin) validatePassword(password []byte) error {
	return bcrypt.CompareHashAndPassword([]byte(u.password), password)
}

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
