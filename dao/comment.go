package dao

import (
	"TikTok/constdef"
	"time"
)

type Comment struct {
	Id         int64  `json:"id,omitempty" gorm:"primaryKey;unique"`
	UserId     int64  `json:"user_id"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
	VideoId    int64  `json:"video_id"`
}

func GetCommentNum(videoId int64) int64 {
	var num int64
	DB.Table(constdef.CommentsTableName).
		Where("video_id = ?", videoId).
		Count(&num)
	return num
}

func GetComments(videoId int64) []Comment {
	var comments []Comment
	DB.Table(constdef.CommentsTableName).
		Where("video_id = ?", videoId).
		Find(&comments)
	return comments
}

func CreateComment(userId int64, content string, videoId int64) *Comment {
	comment := &Comment{UserId: userId, Content: content, VideoId: videoId, CreateDate: time.Now().Format("2006-01-02 15:04:05")}
	DB.Table(constdef.CommentsTableName).
		Create(&comment)
	return comment
}

func DeleteCommentById(Id int64) {
	DB.Delete(&Comment{}, Id)
}
