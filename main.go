package main

import (
	loam "lineOAM/LOAM"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.POST("/webhook", loam.ReceiveCallBack)
	server.Run(":8080")

}
