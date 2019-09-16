package models

import (
	"errors"
	"fmt"
	u "sflogapi/utils"

	googleAuthIDTokenVerifier "github.com/futurenda/google-auth-id-token-verifier"
	"github.com/jinzhu/gorm"
)

type key string

//Userkey references the user data in context
const (
	Userkey key = "user"
)

//User is a struct to rep user account
type User struct {
	Firstname   string `json:"first_name,omitempty"`
	Lastname    string `json:"last_name,omitempty"`
	Email       string `json:"email,omitempty"`
	Googletoken string `json:"google_token,omitempty"`
}

//Login authenticate and check google token
func Login(email string) map[string]interface{} {
	var user User
	fmt.Println(fmt.Sprintf("Logging in user %s", email))
	err := GetDB().Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	resp := u.Message(true, "Logged In")
	resp["user"] = user
	return resp
}

//VerifyUser returns an error if the user is not registered on our system
func VerifyUser(email string) error {
	acc := &User{}
	GetDB().Table("users").Where("email = ?", email).First(acc)
	if acc.Email == "" { //User not found!
		return errors.New("User not found")
	}
	return nil
}

//TestDB is for testing the database
func TestDB(email string) {
	user := &User{}
	err := GetDB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		fmt.Println("Error", err)
	}
}

//TestToken takes a google authorisation token and verifies it
func TestToken(tok string) (googleAuthIDTokenVerifier.ClaimSet, error) {
	v := googleAuthIDTokenVerifier.Verifier{}
	aud := "948082053040-r1tead48gksuq902m1g4fo4rsk5qj1tu.apps.googleusercontent.com"
	err := v.VerifyIDToken(tok, []string{
		aud,
	})
	var claimSet *googleAuthIDTokenVerifier.ClaimSet
	if err == nil {
		claimSet, err := googleAuthIDTokenVerifier.Decode(tok)
		if err != nil {
			fmt.Println("Token not verified against google requirements", err)
			return *claimSet, err
		} else {
			return *claimSet, nil
		}
	} else {
		fmt.Println("Token not verified against client account ", err)
		return *claimSet, err
	}

}
