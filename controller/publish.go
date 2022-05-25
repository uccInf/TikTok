package controller

import (
	"TikTok/constdef"
	"TikTok/service"
	"TikTok/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if claim, err := utils.ParseToken(token); claim == nil || err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	claim, _ := utils.ParseToken(token)
	user, _ := service.GetUserByName(claim.UserName)
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join(constdef.StaticLocalPath, finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	playUrl := "http://" + filepath.Join(
		fmt.Sprintf("%s:%d%s", constdef.Ip, constdef.ServerPort, constdef.StaticServerPath),
		finalName,
	)
	service.CreateVideo(user, playUrl, playUrl)

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	if token != "" {
		if claim, err := utils.ParseToken(token); claim != nil && err == nil {
			videos := service.GetPublishedVideosByUserId(claim.UserId)
			c.JSON(http.StatusOK, VideoListResponse{
				Response: Response{
					StatusCode: 0,
				},
				VideoList: videos,
			})
			return
		}
	}
	c.JSON(http.StatusOK, UserInfoResponse{
		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	})

}
