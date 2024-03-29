package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/glorinli/go-contacts/models"
	u "github.com/glorinli/go-contacts/utils"
	"github.com/gorilla/mux"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)

	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
	}

	contact.UserID = user
	resp := contact.Create()

	u.Respond(w, resp)
}

var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetContacts(uint(id))

	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
