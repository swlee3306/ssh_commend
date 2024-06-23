package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartGinServer() {
	fnc := "StartGinServer"
	log.Printf("%s Run", fnc)

	r := gin.Default()

	// 패닉 및 에러 처리를 위해 리커버리 미들웨어 사용
	r.Use(gin.Recovery())

	//ssh-commend 처리 api
	commApi := r.Group("/api/v1/sshCommend")
	{
		commApi.GET("/comm", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "commend set successfully"})
		})
	}

	err := r.Run(":8080")
	if err != nil {
		log.Printf("%s error: %v", fnc, err)
	}

}
