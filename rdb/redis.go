package rdb

import (
	"TikTok/constdef"
	"TikTok/dao"
	"TikTok/logger"
	"encoding/json"
	"github.com/go-redis/redis"
	"time"
)

var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
	pong, err := rdb.Ping().Result()
	if err != nil {
		logger.Fatal(err)
	}
	if pong != "PONG" {
		logger.Fatal("connect redis fail")
	}
	rdb.Set("WhetherUpdate", "WaiteToUpdate", 5*time.Minute)
	UpdateData()
}

func UpdateData() []dao.Video {
	rdb.LTrim("LatestVideos", 1, 0)
	latestVideos := dao.GetLatestVideos()
	for i := 0; i < len(latestVideos); i++ {
		AddVideo(&latestVideos[i])
	}
	return latestVideos
}

func GetRedisClient() *redis.Client {
	return rdb
}

func UpdateLatestVideos(video *dao.Video) {
	length := rdb.LLen("LatestVideos").Val()
	if length == constdef.MaxFeedVideoNum {
		RemoveVideo()
	}
	AddVideo(video)
}

func AddVideo(video *dao.Video) {
	videoJson, _ := json.Marshal(*video)
	rdb.RPush("LatestVideos", videoJson)
}

func RemoveVideo() {
	rdb.LPop("LatestVideos")
}

func GetVideoList() []dao.Video {
	if _, err := rdb.Get("WhetherUpdate").Result(); err == nil {
		return UpdateData()
	}
	ls := rdb.LRange("LatestVideos", 0, -1).Val()
	length := rdb.LLen("LatestVideos").Val()
	var videos = make([]dao.Video, length)
	var i = 0
	for _, video := range ls {
		json.Unmarshal([]byte(video), &videos[i])
		i++
	}
	return videos
}

func CheckToken(token string) bool {
	if _, err := rdb.Get(token).Result(); err == nil {
		return true
	}
	return false
}
