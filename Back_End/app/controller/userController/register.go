package userController

import (
	"backend/app/services/userService"
	"backend/app/utils"
	"backend/app/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterData struct {
	Username string `json:"username"        binding:"required"`
	Name     string `json:"name"            binding:"required"`
	Password string `json:"password"        binding:"required"`
	UserType int    `json:"user_type"       binding:"required"`
}

// 接收参数
func Register(c *gin.Context) {
	var data RegisterData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	//判断账号是否已经存在
	err = userService.CheckUsername(data.Username)
	if err == nil {
		utils.JsonErrorResponse(c, 200505, "账号已被注册")
		return
	} else if err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//哈希加密密码
	hash,err:=userService.HashPassword(data.Password)
	if err!=nil{
		utils.JsonErrorResponse(c,200602,"加密失败")
		return
	}
	data.Password=hash
	//实现注册
	err = userService.Register(model.User{
		Username: data.Username,
		Name:     data.Name,
		Password: hash,
		UserType: data.UserType,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
