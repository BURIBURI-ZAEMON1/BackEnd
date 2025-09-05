package postService

import (
	"backend/app/model"
	"backend/conf/database"
	"errors"
)

func CreatePost(post model.Post) error {
	result := database.DB.Create(&post)
	return result.Error
}
func GetPostByID(id int) (*model.Post, error) {
	var post model.Post
	result := database.DB.First(&post, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil
}
func GetAllPosts() ([]model.Post, error) {
	var posts []model.Post
	result := database.DB.Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}
func UpdatePost(post model.Post) error {
	result := database.DB.Save(&post)
	return result.Error
}
func DeletePost(id int) error {
	result := database.DB.Delete(&model.Post{}, id)
	return result.Error
}
func ReportPost(ID int, MyID int, reason string) (error) {
	var post model.Post
	result := database.DB.First(&post, ID)
	if result.Error != nil {
		return result.Error
	}
	report := model.Report{
		UserID:  MyID,
		PostID:  post.ID,
		Username: post.Username,
		Content: post.Content,
		Reason:  reason,
		Status:  0,
	}
	reportResult := database.DB.Create(&report)
	if reportResult.Error != nil {
		return reportResult.Error
	}

	return  nil
}
func GetReport(userID int) ([]model.Report, error) {
	userType, err := GetUserType(userID)
	if err != nil {
		return nil, err
	}
	if userType != 2 {
		return nil, errors.New("权限不足")
	}
	var report []model.Report
	result := database.DB.Where("status=?", 0).Find(&report)
	if result.Error != nil {
		return nil, result.Error
	}
	return report, nil
}
func UpdateReportStatus(reportID int, newStatus int) error {
	result := database.DB.Model(&model.Report{}).
		Where("id = ?", reportID).
		Update("status", newStatus)
	return result.Error
}
func GetUserType(userID int) (int, error) {
	var userType int
	result := database.DB.Model(&model.User{}).
		Select("user_type").
		Where("id = ?", userID).
		Scan(&userType)
	if result.Error != nil {
		return 0, result.Error
	}

	return userType, nil
}
func GetPostID(ReportID int) (int, error) {
	var PostID int
	result := database.DB.Model(&model.Report{}).
		Select("post_id").
		Where("id = ?", ReportID).
		Scan(&PostID)
	if result.Error != nil {
		return 0, result.Error
	}
	return PostID, nil
}
func ApprovalReport(userID int, reportID int, approval int) error {
	userType, err := GetUserType(userID)
	if err != nil {
		return err
	}
	if userType != 2 {
		return errors.New("权限不足")
	}
	switch approval {
	case 1:
		postID, err := GetPostID(reportID)
		if err != nil {
			return err
		}
		result := database.DB.Delete(&model.Post{}, postID)
		if result.Error != nil {
			return result.Error
		}
		return UpdateReportStatus(reportID, 1)
	case 2:
		return UpdateReportStatus(reportID, 2)
	default:
		return errors.New("无效操作")
	}
}
func GetReportResult(userID int) ([]model.Report, error) {
	var report []model.Report
	result := database.DB.Where("user_id=?", userID).Find(&report)
	if result.Error != nil {
		return report, result.Error
	}
	return report, nil
}
func GetUsernameByID(userID int) (string, error) {
    var user model.User
    result := database.DB.
        Select("username").
        Where("id = ?", userID).
        First(&user)
    
    if result.Error != nil {
        return "",result.Error
    }
    return user.Username, nil
}