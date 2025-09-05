package router

import (
	adminInterface "backend/app/controller/admininterface"
	studentInterface "backend/app/controller/studentinterface"
	"backend/app/controller/userController"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	const pre = "/api"
	api := r.Group(pre)
	api.POST("/user/login", userController.Login)
	api.POST("/user/reg", userController.Register)
	api.POST("/student/post", studentInterface.Publish)
	api.PUT("/student/post", studentInterface.Updata)
	api.DELETE("/student/post", studentInterface.Delete)
	api.GET("/student/post", studentInterface.GetPosts)
	api.POST("/student/report-post", studentInterface.Report)
	api.GET("/student/report-post", studentInterface.CheckReport)
	api.GET("/admin/report", adminInterface.GetAllReport)
	api.POST("/admin/report", adminInterface.ApprovalAllReport)
}
