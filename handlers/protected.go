package handlers

import (
	"net/http"

	"github.com/amankapoor/basic-authentication-using-cookies/common"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret"))

func Protected(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("Aman")
	if err == http.ErrNoCookie {
		w.Write([]byte("Cookie not set"))
		return
	}

	common.RenderTemplate(w, "templates/protected.html", cookie)
}
