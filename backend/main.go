package main

import (
	"github.com/gin-gonic/gin"
	"github.com/palledad/dManga/router"
)

func main() {
	r := gin.Default()
	router.NewRouter(r)
	_ = r.Run()
}
