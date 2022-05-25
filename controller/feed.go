package controller

import (
	"TikTok/dao"
	"TikTok/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []dao.Video `json:"video_list,omitempty"`
	NextTime  int64       `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")
	if user, exist := usersLoginInfo[token]; exist {
		RecommendVideos := service.GetPublishedVideosByUserId(user.Id)
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0},
			VideoList: RecommendVideos,
			NextTime:  time.Now().Unix(),
		})
	} else {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0},
			VideoList: DemoVideos,
			NextTime:  time.Now().Unix(),
		})
	}
}
