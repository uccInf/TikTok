package controller

import (
	"TikTok/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfoResponse struct {
	Response
	User dao.User `json:"user"`
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
