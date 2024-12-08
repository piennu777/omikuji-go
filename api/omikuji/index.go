package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LotteryResult struct {
	Prize string `json:"prize"`
}

var prizes = []string{
	"大当たり",
	"中当たり",
	"小当たり",
	"大外れ",
	"中外れ",
	"小外れ",
	"帰れ",
}

func drawLottery(c *gin.Context) {
	result := LotteryResult{
		Prize: prizes[rand.Intn(len(prizes))],
	}
	c.JSON(http.StatusOK, result)
}

func main() {
	r := gin.Default()

	r.POST("", drawLottery)
	r.Run(":8080")
}
