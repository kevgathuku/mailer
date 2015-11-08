package controller

import (
	"github.com/kn9ts/frodo"
	"net/http"
)

// Home plays an example of a controller
type Home struct {
	Frodo.Controller
}

func (h *Home) Index(w http.ResponseWriter, r *Frodo.Request) {
	name := r.FormValue("name")
	w.Write([]byte(name))
}
