package studentInterface

import (
	"backend/app/model"
	"backend/app/services/postService"
	"backend/app/utils"

	"github.com/gin-gonic/gin"
)

// 发布帖子
type postdata struct {
	Content string `json:"content" binding:"required"`
	UserID  int    `json:"user_id" binding:"required"`
}

func Publish(c *gin.Context) {
	uid, ok := c.Get("userID")
	if !ok {
		utils.JsonErrorResponse(c, 200401, "未登录")
		return
	}
	var post postdata
	post.UserID=uid.(int)
	err := c.ShouldBindJSON(&post)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	var username string
	username, err = postService.GetUsernameByID(post.UserID)
	if err != nil {
		utils.JsonErrorResponse(c, 500601, "获取用户名失败")
	}
	newPost := model.Post{
		UserID:   post.UserID,
		Content:  post.Content,
		Username: username,
	}
	err = postService.CreatePost(newPost)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "发布帖子失败")
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

// 更新帖子
type updatePostData struct {
	UserID  int    `json:"user_id" binding:"required"`
	ID      int    `json:"post_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func Update(c *gin.Context) {
	uid, ok := c.Get("userID")
	if !ok {
		utils.JsonErrorResponse(c, 200401, "未登录")
		return
	}
	var uppost updatePostData
	uppost.UserID=uid.(int)
	err := c.ShouldBindJSON(&uppost)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	originalPost, err := postService.GetPostByID(int(uppost.ID))
	if err != nil {
		utils.JsonErrorResponse(c, 200508, "帖子不存在")
		return
	}
	if originalPost.UserID != uppost.UserID {
		utils.JsonErrorResponse(c, 200509, "无权")
		return
	}
	updatedPost := model.Post{
		ID:        uppost.ID,
		UserID:    originalPost.UserID,
		Username:  originalPost.Username,
		Content:   uppost.Content,
		CreatedAt: originalPost.CreatedAt,
	}
	err = postService.UpdatePost(updatedPost)
	if err != nil {
		utils.JsonErrorResponse(c, 500510, "更新帖子失败")
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

// 删除帖子
type deletePostData struct {
	UserID int `form:"user_id" binding:"required"`
	PostID int `form:"post_id" binding:"required"`
}

func Delete(c *gin.Context) {
	uid, ok := c.Get("userID")
	if !ok {
		utils.JsonErrorResponse(c, 200401, "未登录")
		return
	}
	var depost deletePostData
	depost.UserID=uid.(int)
	err := c.ShouldBind(&depost)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	originalPost, err := postService.GetPostByID(depost.PostID)
	if err != nil {
		utils.JsonErrorResponse(c, 200508, "帖子不存在")
		return
	}
	if originalPost.UserID != depost.UserID {
		utils.JsonErrorResponse(c, 200509, "无权")
		return
	}
	err = postService.DeletePost(depost.PostID)
	if err != nil {
		utils.JsonErrorResponse(c, 200511, "删除帖子失败")
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

// 举报帖子
type reportPostData struct {
	PostID int    `json:"post_id" binding:"required"`
	MyID   int    `json:"user_id" binding:"required"`
	Reason string `json:"reason" binding:"required"`
}

func Report(c *gin.Context) {
	uid, ok := c.Get("userID")
	if !ok {
		utils.JsonErrorResponse(c, 200401, "未登录")
		return
	}
	var repost reportPostData
	repost.MyID = uid.(int)
	err := c.ShouldBindJSON(&repost)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	err = postService.ReportPost(repost.PostID, repost.MyID, repost.Reason)
	if err != nil {
		utils.JsonErrorResponse(c, 200512, "举报失败")
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

// 获取所有帖子
type getPost struct {
	PostList []model.Post `json:"post_list"`
}

func GetPosts(c *gin.Context) {
	postlist, err := postService.GetAllPosts()
	if err != nil {
		utils.JsonErrorResponse(c, 200513, "获取帖子列表失败")
		return
	}
	utils.JsonSuccessResponse(c, getPost{PostList: postlist})
}

// 分页获取帖子
type pageData struct {
	Page int `form:"page" binding:"required"`
}
type getPostsWithPagination struct {
	PostListstWithPagination []model.Post `json:"post_lists"`
	PostNum                  int          `json:"post_num"`
	PostPage                 int          `json:"post_page"`
}

const pagesize = 10

func GetPostsWithPagination(c *gin.Context) {
	var pagedata pageData
	err := c.ShouldBind(&pagedata)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	postlist, total, err := postService.GetPostWithPagination(pagedata.Page, pagesize)
	if err != nil {
		utils.JsonErrorResponse(c, 200513, "获取帖子列表失败")
		return
	}
	postPage := (total + pagesize - 1) / pagesize
	utils.JsonSuccessResponse(c, getPostsWithPagination{PostListstWithPagination: postlist, PostNum: int(total), PostPage: int(postPage)})
}

// 查看举报审批
type checkData struct {
	UserID int `form:"user_id" binding:"required"`
}
type checkReport struct {
	ReportList []model.Report `json:"report_list"`
}

func CheckReport(c *gin.Context) {
	uid, ok := c.Get("userID")
	if !ok {
		utils.JsonErrorResponse(c, 200401, "未登录")
		return
	}
	var chechdata checkData
	chechdata.UserID=uid.(int)
	err := c.ShouldBind(&chechdata)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	result, err := postService.GetReportResult(chechdata.UserID)
	if err != nil {
		utils.JsonErrorResponse(c, 200516, "获取失败")
		return
	}
	utils.JsonSuccessResponse(c, checkReport{ReportList: result})
}
