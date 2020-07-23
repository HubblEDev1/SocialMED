package main

import (
	"log"
	"github.com/edwin1732/SocialMED/handlers"
	"github.com/edwin1732/SocialMED/bd"
)
//Puerta de entrada
func main() {
	if bd.CheckConnection()==0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manager()

}
