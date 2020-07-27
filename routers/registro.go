package routers

import (
	"encoding/json"
	"net/http"
	"github.com/edwin1732/SocialMED/bd"
	"github.com/edwin1732/SocialMED/models"
)
//Funcion para crear en la base de datos el registro de usuario
func Registro(w http.ResponseWriter, r *http.Request){
	var t models.User
	err:=json.NewDecoder(r.Body).Decode(&t)
	if err != nil{
		http.Error(w, "Error en los datos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0{
		http.Error(w,"El Email es requerido", 400)
		return
	}
	if len(t.Password)<6{
		http.Error(w,"La contraseña debe ser mayor a 6 carácteres", 400)
		return
	}

	_, found, _ := bd.UserIndatabase(t.Email) //Válida que no existan usuarios repetidos
	if found==true {
		http.Error(w,"El Email ya existe en la Base de datos",400)
		return
	}
	_, status, err:=bd.InsertRegister(t)
	if err!=nil {
		http.Error(w,"Ocurrió un error al realizar el registro"+err.Error(),400)
		return
	}

	if status == false{
		http.Error(w,"No se realizó el registro de usuario",404)
		return
	}

	w.WriteHeader(http.StatusCreated)
	
}