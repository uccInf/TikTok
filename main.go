package main

import (
	"TikTok/rdb"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	initRouter(r)

	rdb.InitRedis()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
