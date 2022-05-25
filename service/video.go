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
