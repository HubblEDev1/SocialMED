package bd

import (
	"context"
	"time"
	"github.com/edwin1732/SocialMED/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Inserto registro
func InsertRegister(u models.User)(string, bool, error){
	ctx, cancel:=context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("SocialMED")
	col:=db.Collection("user")

	u.Password, _= EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err!=nil {
		return "",false,err
	}

	/************************************************/
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}