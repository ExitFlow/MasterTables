package controller

import (
	"log"
	"net/http"

	"github.com/josephspurrier/gowebapp/app/model"
	"github.com/josephspurrier/gowebapp/app/shared/passhash"
	"github.com/josephspurrier/gowebapp/app/shared/recaptcha"
	"github.com/josephspurrier/gowebapp/app/shared/session"
	"github.com/josephspurrier/gowebapp/app/shared/view"

	"github.com/josephspurrier/csrfbanana"
)

// RegisterGET displays the register page
func RegisterGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)

	// Display the view
	v := view.New(r)
	v.Name = "register/register"
	v.Vars["token"] = csrfbanana.Token(w, r, sess)
	// Refill any form fields
	view.Repopulate([]string{"first_name", "last_name", "email"}, r.Form, v.Vars)
	v.Render(w)
}

// RegisterPOST handles the registration form submission
func RegisterPOST(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)

	// Prevent brute force login attempts by not hitting MySQL and pretending like it was invalid :-)
	if sess.Values["register_attempt"] != nil && sess.Values["register_attempt"].(int) >= 5 {
		log.Println("Brute force register prevented")
		http.Redirect(w, r, "/register", http.StatusFound)
		return
	}

	// Validate with required fields (campos obligatorios)
	if validate, missingField := view.Validate(r, []string{"first_name", "last_name", "email", "password"}); !validate {
		sess.AddFlash(view.Flash{Message: "Field missing: " + missingField, Class: view.FlashError})
		sess.Save(r, w)
		RegisterGET(w, r)
		return
	}

	// Validate with Google reCAPTCHA
	if !recaptcha.Verified(r) {
		sess.AddFlash(view.Flash{Message: "reCAPTCHA invalid!", Class: view.FlashError})
		sess.Save(r, w)
		RegisterGET(w, r)
		return
	}

	// Get form values
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	email := r.FormValue("email")
	password, errp := passhash.HashString(r.FormValue("password"))

	// If password hashing failed
	if errp != nil {
		log.Println(errp)
		sess.AddFlash(view.Flash{Message: "An error occurred on the server. Please try again later.", Class: view.FlashError})
		sess.Save(r, w)
		http.Redirect(w, r, "/register", http.StatusFound)
		return
	}

	// Get database result
	_, err := model.UserByEmail(email)

	if err == model.ErrNoResult { // If success (no user exists with that email)
		ex := model.UserCreate(firstName, lastName, email, password)
		// Will only error if there is a problem with the query
		if ex != nil {
			log.Println(ex)
			sess.AddFlash(view.Flash{Message: "An error occurred on the server. Please try again later.", Class: view.FlashError})
			sess.Save(r, w)
		} else {
			sess.AddFlash(view.Flash{Message: "Account created successfully for: " + email, Class: view.FlashSuccess})
			sess.Save(r, w)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
	} else if err != nil { // Catch all other errors
		log.Println(err)
		sess.AddFlash(view.Flash{Message: "An error occurred on the server. Please try again later.", Class: view.FlashError})
		sess.Save(r, w)
	} else { // Else the user already exists
		sess.AddFlash(view.Flash{Message: "Account already exists for: " + email, Class: view.FlashError})
		sess.Save(r, w)
	}

	// Display the page
	RegisterGET(w, r)
}
