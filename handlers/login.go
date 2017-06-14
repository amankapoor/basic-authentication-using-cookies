package handlers

import (
	"net/http"
	"time"

	"log"
)

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")

		if email == "aman.sk95@gmail.com" && password != "" {

			cookie := &http.Cookie{
				Name:    "Aman",
				Value:   "Kapoor",
				Expires: time.Now().Add(2 * time.Hour),
			}
			http.SetCookie(w, cookie)
			log.Printf("&cookie gives: %v", &cookie)
			log.Printf("*cookie gives: %v", *cookie)
			log.Printf("cookie gives: %v", cookie)

			http.Redirect(w, r, "http://localhost:8080/protected", http.StatusSeeOther)
		}
	}
}
