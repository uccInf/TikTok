package controller

import (
	"TikTok/dao"
	"TikTok/service"
	"TikTok/utils"
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
	latestVideos := service.GetLatestVideos(0)
	token := c.Query("token")
	if token != "" {
		if claim, err := utils.ParseToken(token); claim != nil && err == nil {
			user, _ := service.GetUserByName(claim.UserName)
			for i := 0; i < len(latestVideos); i++ {
				if service.CheckIsFavorite(latestVideos[i].VideoId, user) {
					latestVideos[i].IsFavorite = true
				}
			}
		}
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: latestVideos,
		NextTime:  time.Now().Unix(),
	})
}
