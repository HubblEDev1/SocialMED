package bd

import (
    "context"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
//Las variables con mayuscula indican que son exportadas
var MongoCN=ConnectDB()
var clientOptions=options.Client().ApplyURI("mongodb+srv://Edwin:Im183.ed@socialmed.jgug0.mongodb.net/<dbname>?retryWrites=true&w=majority")

func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err=client.Ping(context.TODO(),nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión exitosa con la DB")
	return client
}

func CheckConnection()int{
	err:=MongoCN.Ping(context.TODO(),nil)
	if err != nil {
		return 0
	}
	return 1

}