package controller

import "TikTok/dao"

var DemoVideos = []dao.Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "http://192.168.1.25:8080/static/bear.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 80,
		CommentCount:  25,
		IsFavorite:    false,
	},
}

var DemoComments = []dao.Comment{
	{
		Id: 1,

		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = dao.User{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
