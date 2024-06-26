package ssh

import (
	"log"
	"ssh_commend/internal/sysenv"
	"time"

	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
)

func ClientConnection(vm *sysenv.CommendReq) ([]byte, error) {

	fnc := "ClientConnection"
	var err error

	log.Printf("%s: start", fnc)

	auth := goph.Password(vm.Pwd)

	// SSH 키보드 인터랙티브 콜백 함수 정의
	interactiveCallback := func(user, instruction string, questions []string, echos []bool) ([]string, error) {
		answers := make([]string, len(questions))
		for i := range questions {
			// 모든 질문에 대해 같은 패스워드를 반환
			answers[i] = vm.Pwd
		}
		return answers, nil
	}

	client, err := goph.NewConn(&goph.Config{
		User: vm.User,
		Addr: vm.HostIp,
		Port: vm.Port,
		Auth: []ssh.AuthMethod{
			ssh.KeyboardInteractive(interactiveCallback),
		},
		Timeout:  20 * time.Second,
		Callback: ssh.InsecureIgnoreHostKey(),
	})

	if err != nil {
		client, err = goph.NewConn(&goph.Config{
			User:     vm.User,
			Addr:     vm.HostIp,
			Port:     vm.Port,
			Auth:     auth,
			Timeout:  20 * time.Second,
			Callback: ssh.InsecureIgnoreHostKey(),
		})
		if err != nil {
			log.Printf("%s: new client fail: %s", fnc, err)
			return nil, err
		}
	}

	defer client.Close()

	output, err := client.Run(vm.Cmd)
	if err != nil {
		log.Printf("%s: 명령 실행 실패: %s", fnc, err)
		return nil, err
	} else {
		log.Printf("%s: 명령 실행 성공!!!", fnc)
		log.Printf("명령 실행 결과:\n%s\n", string(output))
		return output, nil
	}
}
