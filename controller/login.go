package controller

import (
	"TikTok/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password
	if user, err := service.Login(username, password); err == nil {
		usersLoginInfo[token] = user
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}
}
