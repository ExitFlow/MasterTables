package controller

import (
	"fmt"
	"net/http"

	"github.com/josephspurrier/csrfbanana"
	"github.com/josephspurrier/gowebapp/app/model"
	"github.com/josephspurrier/gowebapp/app/shared/session"
	"github.com/josephspurrier/gowebapp/app/shared/view"
)

func MantenedorTablaGET(w http.ResponseWriter, r *http.Request) {

	sess := session.Instance(r)

	IdtablaM := fmt.Sprintf("%s", sess.Values["M-IdTabla"])

	if IdtablaM == "%!s(<nil>)" {

		IdtablaM = ""
	}
	opciones := session.GetOpts()
	v := view.New(r)

	ListaTablas, _ := model.ListaTablas(IdtablaM)
	CargaCbxListaTabla, _ := model.CargaCbxListaTabla()

	v.Name = "mantenedor/mantenedorTabla"
	v.Vars["token"] = csrfbanana.Token(w, r, sess)
	v.Vars["ListaTablas"] = ListaTablas
	v.Vars["CargaCbxListaTabla"] = CargaCbxListaTabla
	v.Vars["first_name"] = sess.Values["first_name"]
	v.Vars["opciones"] = opciones
	v.Render(w)
}

func MantenedorTablaPOST(w http.ResponseWriter, r *http.Request) {

	sess := session.Instance(r)

	opciones := session.GetOpts()
	v := view.New(r)

	sess.Values["M-IdTabla"] = r.FormValue("IdTabla")

	v.Vars["opciones"] = opciones

	MantenedorTablaGET(w, r)

}

func LimpiarFiltroListaTablaGET(w http.ResponseWriter, r *http.Request) {

	sess := session.Instance(r)

	sess.Values["M-IdTabla"] = ""

	sess.Save(r, w)
	http.Redirect(w, r, "/mantenedor/mantenedorTabla", http.StatusFound)

}
