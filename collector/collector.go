package collector

import (
	"log"
	"ssh_commend/internal/ssh/disk"
	"ssh_commend/internal/sysenv"
	"time"

	lsapi "ssh_commend/internal/lslinker/lsapi"
	lsmd "ssh_commend/internal/lslinker/lsmd"

	"bitbucket.org/okestrolab/baton-ao-sdk/btoutil"
)

func RunResourceSchedulers(vm *sysenv.VmEnv) {
	fnc := "RunResourceSchedulers"

	log.Printf("Collector Start!")
	log.Printf("VM UUID : %s", vm.UUID)
	log.Printf("VM PRV_ID : %d", vm.PrvId)

	var cmq sysenv.CommendReq
	var ls lsapi.Logstash

	cmq.HostIp = vm.Cfg.Vmattr.HostIp
	cmq.User = vm.Cfg.Vmattr.Id
	cmq.Pwd = vm.Cfg.Vmattr.Pwd
	cmq.Port = vm.Cfg.Vmattr.Port

	datas, err := disk.GetDisk(&cmq)

	if err != nil {
		log.Printf("%s: GetDisk failed: %s", fnc, err.Error())
		return
	}

	rootData := &sysenv.DiskData{}

	for i := range datas {
		if datas[i] == nil {
			log.Printf("datas[%d]가 nil입니다.", i)
			continue
		}
			if datas[i].AvailableSize == 0 {
			rootData.AvailableSize += 0
		}
		rootData.AvailableSize += datas[i].AvailableSize

		if datas[i].TotalSize == 0 {
			rootData.TotalSize += 0
		}
		rootData.TotalSize += datas[i].TotalSize
	}

	for i := 1; i < len(datas); i++ {
		// send
		var sndmsg lsmd.OutputMetricData
		{
			p := &sndmsg

			p.Agent.Type = "metric_collector"
			p.Agent.Name = "ssh-collector"
			p.Agent.LocId = btoutil.ToString(vm.PrvId)

			p.Service.Adapter = "vm"
			p.Service.Type = "openstack"

			p.CollectorTime = time.Now().UnixMilli()

			{
				if datas[i].Uuid == "" {
					p.Basic.Vm.VmId = vm.UUID
				} else {
					p.Basic.Vm.VmId = datas[i].Uuid
				}
				p.Basic.Vm.Domain = "default"
				if vm.Cfg.Vmattr.State == 1 {
					p.Basic.Vm.PowerState = "on"
				} else {
					p.Basic.Vm.PowerState = "off"
				}
				p.Basic.Vm.State = vm.Cfg.Vmattr.State
				p.Basic.System.Filesystem.All.Available = rootData.AvailableSize
				p.Basic.System.Filesystem.All.Total = rootData.TotalSize

				p.Basic.System.Filesystem.Each.Name = datas[i].FilesystemName
				p.Basic.System.Filesystem.Each.Available = datas[i].AvailableSize
				p.Basic.System.Filesystem.Each.Total = datas[i].TotalSize
			}
		}

		ls.Init(vm.Cfg.Logstash.Urls)

		if err = ls.Send(sndmsg); err != nil {
			log.Printf("%s: send failed: %s", fnc, err.Error())
			return
		} else {
			log.Printf("%s: logstash send sucess: %s", fnc, vm.Name)
		}
	}
}
