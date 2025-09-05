package userController

import (
	"backend/app/services/userService"
	"backend/app/utils"
	"backend/app/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type Logdata struct {
	id       int
	typecode int
}

// 接收参数
func Login(c *gin.Context) {
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	//判断用户是否存在
	err = userService.CheckUsername(data.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200502, "用户不存在")
		} else {
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}
	//获取用户信息
	var user *model.User
	user, err = userService.Getuser(data.Username)
	if err != nil {

		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//判断密码是否正确
	flag1 := userService.ComparePwd(data.Password, user.Password)
	if !flag1 {
		utils.JsonErrorResponse(c, 200504, "密码错误")
		return
	}
	//登录成功
	logdata := Logdata{
		id:       int(user.ID),
		typecode: user.UserType,
	}
	utils.JsonSuccessResponse(c, logdata)
}
