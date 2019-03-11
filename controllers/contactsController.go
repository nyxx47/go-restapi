package controllers

import (
	"encoding/json"
	"net/http"
	"restapi/models"
	u "restapi/utils"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	contact.UserId = user
	res := contact.Create()
	u.Respond(w, res)
}

var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetContacts(id)
	res := u.Message(true, "success")
	res["data"] = data
	u.Respond(w, res)
}
