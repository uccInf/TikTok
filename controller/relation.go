package controller

import (
	"TikTok/dao"
	"TikTok/service"
	"TikTok/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserListResponse struct {
	Response
	UserList []dao.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")

	if token != "" {
		if claim, err := utils.ParseToken(token); claim != nil && err == nil {
			c.JSON(http.StatusOK, Response{StatusCode: 0})
			return
		}
	}

	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})

}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	token := c.Query("token")
	if token != "" {
		if claim, err := utils.ParseToken(token); claim != nil && err == nil {
			user, _ := service.GetUserByName(claim.UserName)
			c.JSON(http.StatusOK, UserListResponse{
				Response: Response{
					StatusCode: 0,
				},
				UserList: user.Follows,
			})
			return
		}
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: -1,
			StatusMsg:  "User doesn't exist",
		},
		UserList: []dao.User{},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	if token != "" {
		if claim, err := utils.ParseToken(token); claim != nil && err == nil {
			user, _ := service.GetUserByName(claim.UserName)
			c.JSON(http.StatusOK, UserListResponse{
				Response: Response{
					StatusCode: 0,
				},
				UserList: user.Followers,
			})
			return
		}
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: -1,
			StatusMsg:  "User doesn't exist",
		},
		UserList: []dao.User{},
	})
}
