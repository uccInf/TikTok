package service

import "TikTok/dao"

func CreateComment(userId int64, content string, videoId int64) *dao.Comment {
	return dao.CreateComment(userId, content, videoId)
}

func DeleteComment(Id int64) {
	dao.DeleteCommentById(Id)
}

func GetCommentNum(videoId int64) int64 {
	return dao.GetCommentNum(videoId)
}

func GetComments(videoId int64) []dao.Comment {
	return dao.GetComments(videoId)
}
