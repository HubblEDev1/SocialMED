package bd 

import (
	"context"
	"time"
	"github.com/edwin1732/SocialMED/models"
	"go.mongodb.org/mongo-driver/bson"
)

func UserIndatabase(email string)(models.User, bool, string){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db:=MongoCN.Database("SocialMED")
	col:=db.Collection("user")

	condition:=bson.M{"email":email}

	var resultado models.User

	err:=col.FindOne(ctx, condition).Decode(&resultado)
	ID:=resultado.ID.Hex()
	if err != nil{
		return resultado,false,ID
	}
	return resultado,true,ID
}