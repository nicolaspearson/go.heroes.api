package models

import (
	"os"
	"strings"

	u "go-hero/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Token : A JWT struct
type Token struct {
	UserID uint
	jwt.StandardClaims
}

// User : A user object struct
type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token" sql:"-"`
}

// Validate : Validates a user's details
func (user *User) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(user.Email, "@") {
		return u.Message(false, "An email address is required"), false
	}

	if len(user.Password) < 6 {
		return u.Message(false, "A password is required"), false
	}

	// Email must be unique
	temp := &User{}

	// Check for errors and duplicate emails
	err := GetDB().Table("users").Where("email = ?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message(false, "This email address is already in use."), false
	}

	return u.Message(false, "Validation passed"), true
}

// Create : Registers a new user
func (user *User) Create() map[string]interface{} {
	if resp, ok := user.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	GetDB().Create(user)

	if user.ID <= 0 {
		return u.Message(false, "Failed to create user, connection error.")
	}

	// Create new JWT token
	tk := &Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	user.Token = tokenString

	// Delete password
	user.Password = ""

	response := u.Message(true, "User has been created")
	response["user"] = user
	return response
}

// Login : Creates a JWT for the user if the credentials are valid
func Login(email, password string) map[string]interface{} {
	user := &User{}
	err := GetDB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		// Password does not match
		return u.Message(false, "Invalid login credentials. Please try again")
	}
	// Validated
	user.Password = ""

	// Create a JWT
	tk := &Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	// Store the token in the response
	user.Token = tokenString

	resp := u.Message(true, "Logged In")
	resp["user"] = user
	return resp
}

// GetUser : Fetches a user from the database
func GetUser(id uint) *User {
	user := &User{}
	GetDB().Table("users").Where("id = ?", id).First(user)
	if user.Email == "" {
		// User not found
		return nil
	}

	user.Password = ""
	return user
}
