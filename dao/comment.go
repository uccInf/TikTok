package dao

type Comment struct {
	Id         int64  `json:"id,omitempty" gorm:"primaryKey;unique"`
	User       User   `json:"user" gorm:"foreignKey:Id"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
	Video      Video  `gorm:"foreignKey:Id"`
}
