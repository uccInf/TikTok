package dao

type Video struct {
	Id            int64  `json:"id,omitempty" gorm:"primaryKey;unique"`
	Author        User   `json:"author" gorm:"foreignKey:Id"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}
