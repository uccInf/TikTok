package controller

import (
	"TikTok/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	content := c.Query("comment_text")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	fmt.Println(videoId, token, content)
	if user, exist := usersLoginInfo[token]; exist {
		fmt.Println(videoId, token, content, user)
		service.CreateComment(user.Id, content, videoId)
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	comments := service.GetComments(videoId)
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: comments,
	})
}
