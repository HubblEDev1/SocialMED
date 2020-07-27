package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID 				primitive.ObjectID 	`bson:"_id,omitempty" json:"id"`
	Nombre 			string			    `bson:"nombre" json:"nombre,omitempty"`
	Apellidos 		string  			`bson:"apellidos" json:"apellidos,omitempty"`
	FechaNacimiento time.Time  			`bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email 			string 				`bson:"email" json:"email"`
	Password 		string 				`bson:"password" json:"password,omitempty"`
	Perfil 			string 				`bson:"perfil" json:"perfil,omitempty"`
	Banner 			string 				`bson:"banner" json:"banner,omitempty"`
	Biografia 		string 				`bson:"biografia" json:"biografia,omitempty"`
	Ubicacion 		string 				`bson:"ubicacion" json:"ubicacion,omitempty"`
	SitioWeb 		string 				`bson:"sitioWeb" json:"sitioWeb,omitempty"`
}