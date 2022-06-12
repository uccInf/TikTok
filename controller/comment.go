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
	actionType := c.Query("action_type")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if service.CheckToken(token) {
		if claim, err := utils.ParseToken(token); claim != nil && err == nil {
			if actionType == "1" {
				content := c.Query("comment_text")
				comment := service.CreateComment(claim.UserId, content, videoId)
				service.AddVideoCommentNum(videoId)
				c.JSON(http.StatusOK,
					CommentActionResponse{
						Response: Response{StatusCode: 0},
						Comment:  *comment,
					})

			} else {
				c.JSON(http.StatusOK, Response{StatusCode: 0})
			}
			return
		}
	}
	c.JSON(http.StatusOK,
		Response{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist or token has been out of date, please relogin",
		})

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
