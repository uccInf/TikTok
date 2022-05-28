package controller

import (
	"TikTok/dao"
)

// var usersLoginInfo = make(map[string]*dao.User)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserInfoResponse struct {
	Response
	User dao.User `json:"user"`
}

type CommentListResponse struct {
	Response
	CommentList []dao.Comment `json:"comment_list,omitempty"`
}

type VideoListResponse struct {
	Response
	VideoList []dao.Video `json:"video_list"`
}

type CommentActionResponse struct {
	Response
	Comment dao.Comment `json:"comment,omitempty"`
}
