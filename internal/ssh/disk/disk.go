package disk

import (
	"log"
	"ssh_commend/internal/ssh"
	"ssh_commend/internal/ssh/cmdline"
	"ssh_commend/internal/sysenv"
	"strconv"
	"strings"
)

func GetDisk(vm *sysenv.CommendReq) ([]*sysenv.DiskData, error) {
	fnc := "GetDisk"
	log.Printf("%s: start", fnc)

	var datas []*sysenv.DiskData

	vm.Cmd = cmdline.Cmdline.DISK_USE

	output, err := ssh.ClientConnection(vm)
	if err != nil {
		log.Printf("%s: Fail output: %s", fnc, err)
		return nil, err
	}

	output_s := string(output)
	lines := strings.Split(output_s, "\n")
	if len(lines) < 2 {
		log.Printf("%s: Unexpected output: %s", fnc, output)
	}

	uuid := lines[0]

	for i := 1; i < len(lines); i++ {
		fields := strings.Fields(lines[i])

		if len(fields) < 5 {
			log.Printf("%s: Unexpected output format: %s", fnc, lines[i])
			continue
		}

		t, _ := strconv.ParseFloat(fields[1], 64)
		u, _ := strconv.ParseFloat(fields[2], 64)
		a, _ := strconv.ParseFloat(fields[3], 64)

		// 전체 용량, 남은 용량, 사용률 추출
		data := sysenv.DiskData{
			FilesystemName: fields[0],
			TotalSize:      t,
			UsedSize:       u,
			AvailableSize:  a,
			UsagePercent:   fields[4],
			Uuid:           uuid,
		}
		datas = append(datas, &data)
	}

	return datas, nil
}
