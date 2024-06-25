package router

import (
	"log"
	"net/http"
	"ssh_commend/internal/sysenv"

	"github.com/gin-gonic/gin"
)

func StartGinServer() {
	fnc := "StartGinServer"
	log.Printf("%s Run", fnc)

	r := gin.Default()

	// 패닉 및 에러 처리를 위해 리커버리 미들웨어 사용
	r.Use(gin.Recovery())

	//ssh-commend 처리 api
	commApi := r.Group("/api/v1")
	{
		commApi.POST("/sshCommend", func(ctx *gin.Context) {
			var err error
			var req sysenv.CommendReq

			if err = ctx.BindJSON(&req); err != nil {
				log.Printf("%s: ctx.BindJSON faild: %s", fnc, err.Error())
				ctx.JSON(http.StatusBadRequest, gin.H{"message": ":" + err.Error()})
				return
			}

			sysenv.SSH_COMMEND <- req

			ctx.JSON(http.StatusOK, gin.H{"message": "commend set successfully"})
		})
	}

	err := r.Run(":8080")
	if err != nil {
		log.Printf("%s error: %v", fnc, err)
	}

}
