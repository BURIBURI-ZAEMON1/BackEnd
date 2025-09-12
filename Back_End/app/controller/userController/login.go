package userController

import (
	"backend/app/middleware"
	"backend/app/services/userService"
	"backend/app/utils"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type Logdata struct {
	ID       int    `json:"user_id"`
	Typecode int    `json:"user_type"`
	Token    string `json:"token"`
}

// 接收参数
func Login(c *gin.Context) {
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	user, err := userService.GetUser(data.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.JsonErrorResponse(c, 200502, "用户不存在")
		} else {
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}

	if err := userService.CompareHash(data.Password, user.Password); err != nil {
		utils.JsonErrorResponse(c, 200504, "密码错误")
		return
	}
	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		utils.JsonErrorResponse(c, 200401, "生成token失败")
		return
	}
	logdata := Logdata{
		ID:       int(user.ID),
		Typecode: user.UserType,
		Token:    token,
	}
	utils.JsonSuccessResponse(c, logdata)
}
