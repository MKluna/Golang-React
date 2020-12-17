package main

import (
	"log"
	"github.com/MKluna/Golang-React/handlers"
	"github.com/MKluna/Golang-React/bd"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin Conexion a la Base de Datos")
		return
	}
	handlers.Manejadores()
}