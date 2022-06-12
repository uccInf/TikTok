package controller

import (
	"TikTok/constdef"
	"TikTok/logger"
	"TikTok/service"
	"TikTok/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"path/filepath"
)

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if claim, err := utils.ParseToken(token); !service.CheckToken(token) || claim == nil || err != nil {
		c.JSON(http.StatusOK,
			Response{
				StatusCode: 1,
				StatusMsg:  "User doesn't exist or token has been out of date, please relogin",
			})
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
	finalName := fmt.Sprintf("%d_%s", user.UserId, filename)
	saveFile := filepath.Join(constdef.StaticLocalPath, "video", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	playUrl := "http://" + filepath.Join(
		fmt.Sprintf("%s:%d%s", constdef.Ip, constdef.ServerPort, constdef.StaticServerPath),
		"video/"+finalName,
	)
	coverUrl := "http://" + filepath.Join(
		fmt.Sprintf("%s:%d%s", constdef.Ip, constdef.ServerPort, constdef.StaticServerPath),
		"image/"+finalName+".jpg",
	)

	title := c.PostForm("title")

	var ffmpegSite string
	if constdef.CurrentOs == constdef.MacOS {
		ffmpegSite = "./utils/ffmpeg/Mac/ffmpeg"
	} else if constdef.CurrentOs == constdef.Windows {
		ffmpegSite = "./utils/ffmpeg/Windows/ffmpeg"
	}

	cmd := exec.Command(
		ffmpegSite, "-i", "./public/video/"+finalName,
		"-vf", "select=eq(n\\, 10)", "-frames", "1",
		"./public/image/"+finalName+".jpg",
	)

	if err = cmd.Run(); err != nil {
		logger.Error(err.Error())
	}

	service.CreateVideo(user, playUrl, coverUrl, title)
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	if service.CheckToken(token) {
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
		Response: Response{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist or token has been out of date, please relogin",
		},
	})

}
