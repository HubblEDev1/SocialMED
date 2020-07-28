package models

//Tiene el token que se devuelve con el log
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`

}