package handlers

import (
	"log"
	"net/http"
	"os"
	"github.com/edwin1732/SocialMED/middlew"
	"github.com/edwin1732/SocialMED/routers"	
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)
//Set del puerto y pongo a escuchar mi server
func Manager(){
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.CheckDB(routers.Registro)).Methods("POST") //Endpoint
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST") //Endpoint
	
	PORT := os.Getenv("PORT")
	if PORT==""{
		PORT="8080"
	}
	handler:=cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT,handler))
}