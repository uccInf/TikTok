package dao

import "TikTok/constdef"

type Video struct {
	Id            int64  `json:"id,omitempty" gorm:"primaryKey;unique"`
	Author        User   `json:"author" gorm:"foreignKey:Id;association_foreignKey:AuthorId"`
	AuthorId      int64  `json:"author_id"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

func GetVideoById(videoId int64) *Video {
	var video Video
	DB.Table(constdef.VideosTableName).
		Where("video_id = ?", videoId).
		Find(&video)
	return &video
}

func CreateVideo(author *User, playUrl string, coverUrl string) {
	video := &Video{Author: *author, AuthorId: author.Id, PlayUrl: playUrl, CoverUrl: coverUrl}
	DB.Table(constdef.VideosTableName).Create(video)
}

func GetPublishedVideosByUserId(userId int64) []Video {
	var videos []Video
	DB.Table("videos").
		Where("author_id = ?", userId).
		Find(&videos)
	return videos
}
