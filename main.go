package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"ssh_commend/internal/ssh"
	"ssh_commend/internal/sysdef"
	"ssh_commend/internal/sysenv"
	"ssh_commend/utils/router"
	"strconv"
	"syscall"

	"bitbucket.org/okestrolab/baton-om-sdk/btoutil"
)

var (
	X_buildDatetime, X_buildRevision, X_buildRevisionShort, X_buildBranch, X_buildTag string
)

func main() {
	fnc := "main"
	err := error(nil)

	log.Printf("%s: 빌드 정보", fnc)
	log.Printf("\t buildDatetime: %s", X_buildDatetime)
	log.Printf("\t buildRevision: %s (%s)", X_buildRevisionShort, X_buildRevision)
	log.Printf("\t buildBranch: %s", X_buildBranch)
	log.Printf("\t buildTag: %s", X_buildTag)

	// 초기화
	{
		// 시간 위치 설정
		//btoutil.SetDefaultTimeZone("UTC")
		btoutil.SetDefaultTimeZone("Asia/Seoul")

		// 디버그 모드 설정
		if val, ok := os.LookupEnv("OKE_DEBUG"); ok && (len(val) > 0) {
			sysenv.Mode.IsDebug, _ = strconv.ParseBool(val)
		}

		// 설정 파일 이름 설정
		if val, ok := os.LookupEnv("BATON_SETTING_FILENAME"); ok && (len(val) > 0) {
			sysdef.ConfFilename = val
		}

		// 설정 파일 로드
		if err = main_LoadYml(sysdef.ConfFilename); err != nil {
			panic(fmt.Sprintf("%s: Cfg load 실패: %s", fnc, err.Error()))
		}

		// 데이터베이스 환경 변수 로드
		if err = main_LoadEnvDb(); err != nil {
			panic(fmt.Sprintf("%s: main_LoadEnvDb 실패: %s", fnc, err.Error()))
		}
	}

	go func() {
		for {
			req := <-sysenv.SSH_COMMEND

			var res sysenv.CommendRes

			output, err := ssh.ClientConnection(&req)

			res.Output = string(output)
			if err != nil {
				res.Err = err.Error()
			} else {
				res.Err = ""
			}

			sysenv.SSH_COMMEND_RES <- res
		}
	}()

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
