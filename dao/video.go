package dao

import (
	"TikTok/constdef"
	"time"
)

type Video struct {
	VideoId       int64  `json:"id,omitempty" gorm:"primaryKey"`
	Author        User   `json:"author" gorm:"foreignkey:AuthorId"`
	AuthorId      int64  `json:"author_id"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	FavoriteUsers []User `gorm:"many2many:user_videos"`
	CreateTime    string `json:"create_time"`
}

func GetVideoById(videoId int64) *Video {
	var video Video
	DB.Table(constdef.VideosTableName).
		Preload("Author").
		Where("video_id = ?", videoId).
		First(&video)
	return &video
}

func CreateVideo(author *User, playUrl string, coverUrl string) {
	video := &Video{
		Author:     *author,
		AuthorId:   author.UserId,
		PlayUrl:    playUrl,
		CoverUrl:   coverUrl,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	DB.Table(constdef.VideosTableName).
		Preload("Author").
		Create(video)
}

func GetPublishedVideosByUserId(userId int64) []Video {
	var videos []Video
	DB.Table(constdef.VideosTableName).
		Preload("Author").
		Where("author_id = ?", userId).
		Find(&videos)
	return videos
}

func AddVideoCommentNum(videoId int64, num int64) {
	DB.Model(Video{}).Where("video_id = ?", videoId).Update("comment_count", num+1)
}

func AddVideoFavoriteNum(videoId int64, num int64) {
	DB.Model(Video{}).Where("video_id = ?", videoId).Update("favorite_count", num+1)
}

func RemoveVideoFavoriteNum(videoId int64, num int64) {
	DB.Model(Video{}).Where("video_id = ?", videoId).Update("favorite_count", num-1)
}

func GetLatestVideos(skip int) []Video {
	var videos []Video
	DB.Table(constdef.VideosTableName).
		Preload("Author").
		Order("create_time desc").
		Limit(30).
		Offset(skip).
		Find(&videos)
	return videos
}
