package controllers

import (
	"encoding/json"
	"net/http"
	"restapi/models"
	u "restapi/utils"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	res := account.Create()
	u.Respond(w, res)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	res := models.Login(account.Email, account.Password)
	u.Respond(w, res)
}
