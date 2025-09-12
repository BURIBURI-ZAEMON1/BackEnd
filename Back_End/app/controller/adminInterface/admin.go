package adminInterface

import (
	"backend/app/model"
	"backend/app/services/postService"
	"backend/app/utils"

	"github.com/gin-gonic/gin"
)

// 获取未审批的被举报帖子
type GetData struct {
	UserID int `form:"user_id" binding:"required"`
}
type getReport struct {
	ReportsList []model.Report `json:"report_list"`
}

func GetAllReport(c *gin.Context) {
	uid, ok := c.Get("userID")
	if !ok {
		utils.JsonErrorResponse(c, 200401, "未登录")
		return
	}
	var getdata GetData
	getdata.UserID=uid.(int)
	err := c.ShouldBind(&getdata)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	reports, err := postService.GetReport(getdata.UserID)
	if err != nil {
		utils.JsonErrorResponse(c, 200514, "获取举报列表失败")
		return
	}
	utils.JsonSuccessResponse(c, getReport{ReportsList:reports})
}

// 审核被举报的帖子
type ApprovalData struct {
	UserID   int `json:"user_id" binding:"required"`
	ReportID int `json:"report_id" binding:"required"`
	Approval int `json:"approval" binding:"required"`
}

func ApprovalAllReport(c *gin.Context) {
	uid, ok := c.Get("userID")
	if !ok {
		utils.JsonErrorResponse(c, 200401, "未登录")
		return
	}
	var approvaldata ApprovalData
	approvaldata.UserID=uid.(int)
	err := c.ShouldBindJSON(&approvaldata)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	err = postService.ApprovalReport(approvaldata.UserID, approvaldata.ReportID, approvaldata.Approval)
	if err != nil {
		utils.JsonErrorResponse(c, 500515, "审核失败")
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
