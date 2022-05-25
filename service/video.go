package service

import (
	"TikTok/dao"
)

func GetVideoByVideoId(videoId int64) *dao.Video {
	return dao.GetVideoById(videoId)
}

func CreateVideo(author *dao.User, playUrl string, coverUrl string) {
	dao.CreateVideo(author, playUrl, coverUrl)
}

func GetPublishedVideosByUserId(userId int64) []dao.Video {
	return dao.GetPublishedVideosByUserId(userId)
}

func AddVideoCommentNum(videoId int64) {
	num := dao.GetVideoById(videoId).CommentCount
	dao.AddVideoCommentNum(videoId, num)
}
