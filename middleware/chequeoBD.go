package middleware

import (
	"net/http"

	"github.com/MKluna/Golang-React/bd"
)

/*ChequeoBD es el middleware que me permite conocer el estado de la base de datos*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Lost connection", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
