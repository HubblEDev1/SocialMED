package checkdb

import (
	"net/http"
	"github.com/edwin1732/SocialMED/bd"
)
//CheckDB es el middlew que me permite conocer el estado de la base de datos
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection()==0{
			http.Error(w, "Connection lost with DB", 500)
			return 
		}
		next.ServeHTTP(w, r)
	}
}