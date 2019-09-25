package admin

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var err = errors.New("Invalid input. Please enter valid email and password")

//Admin is the system user who has complete control on the website
//he can then edit all the contents on the websites
type Admin struct {
	Name        string
	Email       string
	Password    string
	AccessToken string
}

var privateKey = []byte("abcdefghijxyzABCDEopqrFGHIJKLklmnstuvwMN345678OPQRSTUVWXYZ1290")

func hashAndSalt(password []byte) string {
	encrypted, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic("While creating password: " + err.Error())
	}
	return string(encrypted)
}

//isValidPassword will validate password for user
func (u *Admin) isValidPassword(password []byte) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), password); err != nil {
		return false
	}
	return true
}

//GenerateAccessToken will generate access token for user
func (u *Admin) generateAccessToken() {
	token := jwt.New(jwt.SigningMethodHS256)
	var err error
	u.AccessToken, err = token.SignedString(privateKey)
	if err != nil {
		panic(err)
	}
}

//ValidateRequest will validate request and will return token
func ValidateRequest(email string, password string) (*string, error) {
	if email != ADMIN.Email || !ADMIN.isValidPassword([]byte(password)) {
		return nil, err
	}
	ADMIN.generateAccessToken()
	token := ADMIN.AccessToken
	return &token, nil
}

//DestroyAccessToken will be called on signout call
func (u *Admin) DestroyAccessToken() {
	u.AccessToken = ""
}

//RefreshToken will regenerate the token and will send it to client
func RefreshToken() *string {
	ADMIN.generateAccessToken()
	return &ADMIN.AccessToken
}

//IsValidToken will check if the passed token is valid or not
func IsValidToken(token string) bool {
	parsedToken, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})
	if parsedToken.Valid {
		return true
	}
	return false
}
