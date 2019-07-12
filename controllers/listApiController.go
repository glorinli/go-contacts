package controllers

import (
	"net/http"

	u "github.com/glorinli/go-contacts/utils"
)

var ListApi = func(w http.ResponseWriter, h *http.Request) {
	resp := u.Message(true, "success")

	resp["apis"] = map[string]interface{}{
		"listApi":  "/api",
		"register": "/api/user/new",
	}

	u.Respond(w, resp)
}
