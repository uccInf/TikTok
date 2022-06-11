package controller

import (
	"TikTok/dao"
	"TikTok/service"
	"TikTok/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")
	if token != "" {
		if claim, err := utils.ParseToken(token); claim != nil && err == nil {
			user, _ := service.GetUserByName(claim.UserName)
			videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
			video := service.GetVideoByVideoId(videoId)
			if actionType == "1" {
				service.AddFavorite(user, video)
			} else {
				service.RemoveFavorite(user, video)
			}
			c.JSON(http.StatusOK, Response{StatusCode: 0})
			return
		}
	}
	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	if token != "" {
		if claim, err := utils.ParseToken(token); claim != nil && err == nil {
			user, _ := service.GetUserByName(claim.UserName)
			c.JSON(http.StatusOK, VideoListResponse{
				Response: Response{
					StatusCode: 0,
				},
				VideoList: user.FavoriteVideos,
			})
		}
	} else {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "User doesn't exist",
			},
			VideoList: []dao.Video{},
		})
	}
}
