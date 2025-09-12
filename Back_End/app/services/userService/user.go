package userService

import (
	"backend/app/model"
	"backend/conf/database"

	"golang.org/x/crypto/bcrypt"
)
func CheckUsername (username string) error{
	result:=database.DB.Where("username=?",username).First(&model.User{})
	return result.Error
}
func GetUser(username string)(*model.User,error){
	var user model.User
	result:=database.DB.Where("username=?",username).First(&user)
	if result.Error!=nil{
		return nil,result.Error
	}
	return &user,nil
}
func ComparePwd(pwd1 string,pwd2 string)bool{
	return pwd1==pwd2
}
func Register(user model.User)error{
	result:=database.DB.Create(&user)
	return result.Error
}
func HashPassword(password string)(string,error){
	hash,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err!=nil{
		return  "",err
	}
	return string(hash),nil
}
func CompareHash(password,hash string)error{
	err:=bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	return  err
}