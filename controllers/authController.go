package controllers

import (
	"fmt"
	"net/http"
	"sflogapi/models"
	u "sflogapi/utils"

	googleAuthIDTokenVerifier "github.com/futurenda/google-auth-id-token-verifier"
)

//Authenticate sends the request to be authenticated
var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	// account := &models.User{}

	// b, err := ioutil.ReadAll(r.Body)
	// defer r.Body.Close()
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }

	// err = json.Unmarshal(b, &account) //decode the request body into struct and fail if any error occurs
	// if err != nil {
	// 	fmt.Println("Error", err)
	// 	u.Respond(w, u.Message(false, "Failed decoding to user struct"))
	// 	return
	// }

	var claimSet *googleAuthIDTokenVerifier.ClaimSet
	claimSet, ok := u.GetContext(w, r, u.Userkey).(*googleAuthIDTokenVerifier.ClaimSet)
	if !ok {
		fmt.Printf("%T\n", u.GetContext(w, r, u.Userkey))
		u.Respond(w, u.Message(false, "Failed to get user context"))
	}
	email := claimSet.Email
	resp := models.Login(email)
	u.Respond(w, resp)
}
