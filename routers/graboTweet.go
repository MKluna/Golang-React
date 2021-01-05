package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MKluna/Golang-React/bd"
	"github.com/MKluna/Golang-React/models"
)

/*GraboTweet permite grabar el tweet en la base de datos*/
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)
	registro := models.GraboTweet{
		UserID: IDUsuario,
		Menaje: mensaje.Mensaje,
		Fecha:  time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)

	if err != nil {
		http.Error(w, "No se pudo insertar el registro "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se pudo insertar el registro ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
