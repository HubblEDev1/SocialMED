net/http paquete de net

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