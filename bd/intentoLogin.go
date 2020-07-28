package bd

import (
	"github.com/edwin1732/SocialMED/models"
	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string, password string) (models.User, bool) {
	usu, found,_:=UserIndatabase(email)
	if found==false{
		return usu, false
	}
	passwordBytes:=[]byte(password)
	passwordBD:=[]byte(usu.Password)

	err:=bcrypt.CompareHashAndPassword(passwordBD,passwordBytes)
	if err!=nil{
		return usu, false
	}
	return usu, true
}