package osinfo

import (
	"log"
	"ssh_commend/internal/ssh"
	"ssh_commend/internal/ssh/cmdline"
	"ssh_commend/internal/sysenv"
	"strings"
)

func GetOsInfo(vm *sysenv.CommendReq) ([]*sysenv.OsInfo, error) {
	fnc := "GetOsInfo"
	log.Printf("%s: start", fnc)

	var datas []*sysenv.OsInfo

	vm.Cmd = cmdline.Cmdline.OS_KIND

	output, err := ssh.ClientConnection(vm)
	if err != nil {
		log.Printf("%s: Fail output: %s", fnc, err)
		return nil, err
	}

	output_s := string(output)
	lines := strings.Split(output_s, "\n")

	data := sysenv.OsInfo{
		Osname:    lines[0],
		Osversion: lines[1],
		Desc:      lines[3],
	}

	datas = append(datas, &data)

	return datas, nil
}
