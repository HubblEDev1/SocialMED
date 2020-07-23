net/http paquete de net
//En go los contextos son espacios de memoria en los cuales puedo tener  
**********some comands*********
 heroku login

 Create a new Git repository

Initialize a git repository in a new or existing directory

$ cd my-project/
$ git init
$ heroku git:remote -a mysocialmed
Deploy
$ git add .
$ git commit -am "make it better"
$ git push heroku master

Existing Git repository

For existing repositories, simply add the heroku remote

$ heroku git:remote -a mysocialmed

***********Base de datos de Mongo****************
Mongo es una base de datos nosql
****************************************************Cadena de conecion a bas de datos mongo*****************************************************************************************************
mongodb+srv://Edwin:<password>@socialmed.jgug0.mongodb.net/<dbname>?retryWrites=true&w=majority
*************************************Cadena de conexion compass Mongo***************************************************
mongodb+srv://Edwin:<password>@socialmed.jgug0.mongodb.net/test
**********Dependencias********************************
Use in go.mod 

module github.com/edwin1732/SocialMED

go 1.23 <-------------------Or dont build app

require (
)
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options 
go get go.mongodb.org/mongo-driver/bson  
go get go.mongodb.org/mongo-driver/bson/primitive
go get golang.org/x/crypto/bcrypt
go get github.com/gorilla/mux   
go get github.com/rs/cors Nos permite otorgar permisos a nuestra API para ser accedida remotamente
go get github.com/dgrijalva/jwt-go