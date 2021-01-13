package routers

import (
	"net/http"

	"github.com/MKluna/Golang-React/bd"
	"github.com/MKluna/Golang-React/models"
)

/*AltaRelacion realiza el registro de la relacion entre entre usuarios*/
func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion

	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar la relacion "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar la relacion "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
