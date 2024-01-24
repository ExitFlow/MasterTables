package controller

import (
	"net/http"

	"github.com/josephspurrier/csrfbanana"
	"github.com/josephspurrier/gowebapp/app/shared/session"
	"github.com/josephspurrier/gowebapp/app/shared/view"
)

// IndexGET displays the home page
func IndexGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)

	if sess.Values["id"] != nil {
		// Display the view
		v := view.New(r)
		v.Name = "index/auth"
		opciones := session.GetOpts()
		v.Vars["token"] = csrfbanana.Token(w, r, sess)
		v.Vars["opciones"] = opciones
		v.Vars["first_name"] = sess.Values["first_name"]

		v.Render(w)

	} else {
		// Display the view
		v := view.New(r)
		v.Name = "index/anon"
		v.Render(w)
		return
	}
}
