package collector

import (
	"log"
	"ssh_commend/internal/ssh/osinfo"
	"ssh_commend/internal/sysenv"
)

func MetaUpdate(vm *sysenv.VmEnv){
	fnc := "MetaUpdate"

	log.Printf("Collector Start!")
	log.Printf("VM UUID : %s", vm.UUID)
	log.Printf("VM PRV_ID : %d", vm.PrvId)

	var cmq sysenv.CommendReq

	cmq.HostIp = vm.Cfg.Vmattr.HostIp
	cmq.User = vm.Cfg.Vmattr.Id
	cmq.Pwd = vm.Cfg.Vmattr.Pwd
	cmq.Port = vm.Cfg.Vmattr.Port

	datas, err := osinfo.GetOsInfo(&cmq)

	if err != nil {
		log.Printf("%s: GetOsInfo failed: %s", fnc, err.Error())
		return
	}

	rootData := &sysenv.OsInfo{}

	for i := range datas {
		if datas[i] == nil {
			log.Printf("datas[%d]가 nil입니다.", i)
			continue
		}
		if datas[i].Osname == "" {
			rootData.Osname = "ubuntu"
		}
		rootData.Osname = datas[i].Osname
		
		if datas[i].Osversion == "" {
			rootData.Osversion = "latest"
		}
		rootData.Osversion = datas[i].Osversion
	}

	

}