package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/edwin1732/SocialMED/bd"
	"github.com/edwin1732/SocialMED/jwt"
	"github.com/edwin1732/SocialMED/models"
)

func Login(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")
	var t models.User

	err:=json.NewDecoder(r.Body).Decode(&t)
	if err!=nil{
		http.Error(w, "Usuario o contraseña invalida"+err.Error(), 400)
		return
	}
	if len(t.Email)==0{
		http.Error(w, "El Email es requerido", 400)
		return
	}
	documento, existe:=bd.IntentoLogin(t.Email, t.Password)
	if existe==false{
		http.Error(w, "Email y/o Contraseña invalido",400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w,"Ocurrió un error al intentar generar el Token"+err.Error(), 400)
		return
	}
	resp:=models.RespuestaLogin{
		Token : jwtKey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	//Guardar cookies
	expirationTime := time.Now().Add(24*time.Hour)
	http.SetCookie (w, &http.Cookie{
		Name: "token", 
		Value: jwtKey,
		Expires: expirationTime,
	})
}