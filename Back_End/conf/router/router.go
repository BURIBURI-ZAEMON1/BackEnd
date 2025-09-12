package router

import (
	"backend/app/controller/adminInterface"
	"backend/app/controller/studentInterface"
	"backend/app/controller/userController"
	"backend/app/middleware"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	const pre = "/api"
	api := r.Group(pre)
	api.POST("/user/login", userController.Login)
	api.POST("/user/reg", userController.Register)
	auth := api.Group("")
	auth.Use(middleware.JWT())
	{
		auth.POST("/student/post", studentInterface.Publish)
		auth.PUT("/student/post", studentInterface.Update)
		auth.DELETE("/student/post", studentInterface.Delete)
		auth.GET("/student/post", studentInterface.GetPosts)
		auth.POST("/student/report-post", studentInterface.Report)
		auth.GET("/student/report-post", studentInterface.CheckReport)
		auth.GET("/student/postwithpage", studentInterface.GetPostsWithPagination)
		auth.GET("/admin/report", adminInterface.GetAllReport)
		auth.POST("/admin/report", adminInterface.ApprovalAllReport)
	}
}
