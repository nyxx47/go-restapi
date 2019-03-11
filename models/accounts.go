package models

import (
	"os"
	u "restapi/utils"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

func (account *Account) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address required"), false
	}

	if len(account.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	tmpl := &Account{}

	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(tmpl).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if tmpl.Email != "" {
		return u.Message(false, "Email address already in use by another user"), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (account *Account) Create() map[string]interface{} {

	if res, ok := account.Validate(); !ok {
		return res
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)

	account.Password = string(hashedPassword)

	GetDB().Create(account)

	if account.ID <= 0 {
		return u.Message(false, "Failed to create account. connection error.")
	}

	// Buat token baru untuk setiap user yang register
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))

	account.Token = tokenString

	account.Password = ""

	response := u.Message(true, "Account has been created")
	response["account"] = account
	return response
}

func Login(email, password string) map[string]interface{} {

	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email adress not found")
		}

		return u.Message(false, "connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Invalid login credentials. Please try again")
	}

	account.Password = ""

	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	response := u.Message(true, "Logged in")
	response["account"] = account
	return response
}

func GetUser(id uint) *Account {

	users := &Account{}
	GetDB().Table("accounts").Where("id = ?", id).First(users)
	if users.Email == "" {
		return nil
	}

	users.Password = ""
	return users
}
