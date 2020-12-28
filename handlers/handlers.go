package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/MKluna/Golang-React/middleware"
	"github.com/MKluna/Golang-React/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores Set Port , Set Handler and Listen Server*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middleware.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middleware.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middleware.ChequeoBD(middleware.ValidoJWT(routers.VerPerfil))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
