package service

import (
	"TikTok/dao"
	"TikTok/rdb"
)

func GetVideoByVideoId(videoId int64) *dao.Video {
	return dao.GetVideoById(videoId)
}

func CreateVideo(author *dao.User, playUrl string, coverUrl string, title string) {
	video := dao.CreateVideo(author, playUrl, coverUrl, title)
	rdb.UpdateLatestVideos(video)
}

func GetPublishedVideosByUserId(userId int64) []dao.Video {
	return dao.GetPublishedVideosByUserId(userId)
}

func AddVideoCommentNum(videoId int64) {
	num := dao.GetVideoById(videoId).CommentCount
	dao.AddVideoCommentNum(videoId, num)
}

func CheckIsFavorite(videoId int64, user *dao.User) bool {
	for i := 0; i < len(user.FavoriteVideos); i++ {
		if user.FavoriteVideos[i].VideoId == videoId {
			return true
		}
	}
	return false
}

func GetLatestVideos() []dao.Video {
	return rdb.GetVideoList()
}
