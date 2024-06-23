package main

import (
	"log"
	"os"
	"os/signal"
	"ssh_commend/utils/router"
	"sync"
	"syscall"
)

var (
	X_buildDatetime, X_buildRevision, X_buildRevisionShort, X_buildBranch, X_buildTag string
	mutex                                                                             sync.Mutex
)

func main() {
	fnc := "main"
	//err := error(nil)

	log.Printf("%s: 빌드 정보", fnc)
	log.Printf("\t buildDatetime: %s", X_buildDatetime)
	log.Printf("\t buildRevision: %s (%s)", X_buildRevisionShort, X_buildRevision)
	log.Printf("\t buildBranch: %s", X_buildBranch)
	log.Printf("\t buildTag: %s", X_buildTag)

	// Gin 서버를 시작합니다.
	router.StartGinServer()

	// 무한 루프
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for sig := range c {
		log.Printf("인터럽트 발생: %v", sig)

		if sig == syscall.SIGINT || sig == syscall.SIGTERM {
			break
		}
	}
}
