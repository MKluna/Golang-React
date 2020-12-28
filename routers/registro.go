package routers

import (
	"encoding/json"
	"net/http"
	"github.com/MKluna/Golang-React/bd"
	"github.com/MKluna/Golang-React/models"
)
/*Registro es la funcion para crear en la BD el registro de usuario*/
func Registro(w http.ResponseWriter, r *http.Request){
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil{
		http.Error(w, "Los datos recibidos "+ err.Error(),400)
		return
	}

	if len(t.Email)==0{
		http.Error(w, "El email de usuario es requerido",400)
		return
	}

	if len(t.Password)<6{
		http.Error(w, "La contraseña no debe ser menor de seis caracteres",400)
		return
	}

	_,encontrado,_ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		http.Error(w, "Ya existe un usurio registrado con ese email",400)
		return
	}

	_,status,err := bd.InsertoRegistro(t)
	if err !=nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro del usuario :"+err.Error(),400)
		return
	}

	if !status {
		http.Error(w, "No se logro insertar el registro del usuario",400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}