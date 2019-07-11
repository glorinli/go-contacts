package models

import (
	"github.com/dgrijalva/jwt-go"
	u "lens/utils"
	"strings"
	"github.com/jinzhu/gorm"
	"os"
	"golang.org/x/crypto/bcrypt"
)

/*
JWT claims struct
*/
type Token struct {
	UserId uint
	jwt.StandardClaims
}

// a struct to rep user account
type Account struct {
	gorm.Model
	Email string `json: "email"`
	Password string `json:"password"`
	Token string `json:"token";sql:"-"`
}

// validate imcoming user details
func (account *Account) Validate() (map[string] interface{}, bool) {
	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(account.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	temp := &Account{}

	// check for errors and duplicated email
	err := GetDB().Table("accounts").Where("email = ?", account.Email).first(temp).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error, please retry"), false
	}

	if temp.Email != "" {
		return u.Message(false, "Email has been taken"), false
	}

	return u.Message(true, "Checking passed"), true
}

