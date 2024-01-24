package controller

import (
	"net/http"

	"github.com/josephspurrier/gowebapp/app/shared/session"
	"github.com/josephspurrier/gowebapp/app/shared/view"
)

// AboutGET displays the About page
func AboutGET(w http.ResponseWriter, r *http.Request) {
	// Display the view
	v := view.New(r)
	v.Name = "about/about"

	opciones := session.GetOpts()
	v.Vars["opciones"] = opciones

	v.Render(w)
}
