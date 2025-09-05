package userService

import (
	"backend/conf/database"
	"backend/app/model"

)

func CheckUsername (username string) error{
	result:=database.DB.Where("username=?",username).First(&model.User{})
	return result.Error
}
func Getuser(username string)(*model.User,error){
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