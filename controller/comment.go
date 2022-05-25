package controller

import (
	"TikTok/service"
	"TikTok/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	content := c.Query("comment_text")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if token != "" {
		if claim, err := utils.ParseToken(token); claim != nil && err == nil {
			service.CreateComment(claim.UserId, content, videoId)
			service.AddVideoCommentNum(videoId)
			c.JSON(http.StatusOK, Response{StatusCode: 0})
			return
		}

	}
	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})

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
