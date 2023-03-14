package main

import (
	"github.com/gin-gonic/gin"
)

type Danmu struct {
	Rgb     string `json:"rgb"`
	Content string `json:"content"`
}

func main() {
	var danmus [30]Danmu
	e1 := Danmu{Rgb: "#39C5BB", Content: "我超 初音未来"}
	e2 := Danmu{Rgb: "#66CCFF", Content: "我超 洛天依"}
	for i := 0; i <= 15; i++ {
		danmus[i] = e1
	}
	for i := 15; i <= 29; i++ {
		danmus[i] = e2
	}
	r := gin.Default()

	r.GET("/danmu", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": danmus})
	})

	r.Run(":8080")
}
