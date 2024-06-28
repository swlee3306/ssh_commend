package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"ssh_commend/collector"
	"ssh_commend/internal/dblinker"
	"ssh_commend/internal/ssh"
	"ssh_commend/internal/sysdef"
	"ssh_commend/internal/sysenv"
	"ssh_commend/utils/router"
	"strconv"
	"syscall"
	"time"

	"bitbucket.org/okestrolab/baton-ao-sdk/btoutil"
	"gorm.io/gorm"
)

var (
	X_buildDatetime, X_buildRevision, X_buildRevisionShort, X_buildBranch, X_buildTag string
	Db *gorm.DB
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

	Db, err = dblinker.InitDB()
	if err != nil {
		log.Printf("%s DB connection Fail : %s", fnc, err)
	}


	go func() {
		for {
			select {
			case req := <-sysenv.SSH_COMMEND:
				var res sysenv.CommendRes
				output, err := ssh.ClientConnection(&req)
				res.Output = string(output)
				res.User = req.User

				if err != nil {
					res.Err = err.Error()
				} else {
					res.Err = ""
				}
				sysenv.SSH_COMMEND_RES <- res

			case req := <-sysenv.SSH_COMMEND_ALL:
				if req.Status {

					vmlist, err := dblinker.Dbconnection(Db)
					if err != nil {
						log.Printf("%s: Dbconnection 실패: %s", fnc, err.Error())
					}

					var vmlists []*sysenv.CommendRes

					for i := range vmlist {

						var cmq sysenv.CommendReq
						var res sysenv.CommendRes

						cmq.HostIp = vmlist[i].Cfg.Vmattr.HostIp
						cmq.User = vmlist[i].Cfg.Vmattr.Id
						cmq.Pwd = vmlist[i].Cfg.Vmattr.Pwd
						cmq.Port = vmlist[i].Cfg.Vmattr.Port
						cmq.Cmd = req.Cmd

						output, err := ssh.ClientConnection(&cmq)

						res.Output = string(output)
						res.User = cmq.User

						if err != nil {
							res.Err = err.Error()
						} else {
							res.Err = ""
						}

						vmlists = append(vmlists, &res)
					}

					res_e := sysenv.CommendAllRes{
						Vm: vmlists,
					}

					sysenv.SSH_COMMEND_ALL_RES <- res_e
				}
			}
		}
	}()

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			vmlist, err := dblinker.Dbconnection(Db)

			if err != nil {
				log.Printf("Dbconnection 실패: %s", err.Error())
				continue
			}
			for i := range vmlist {
				go func(vm *sysenv.VmEnv) {
					collector.RunResourceSchedulers(vm)
				}(vmlist[i])
			}
		}
	}()

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
