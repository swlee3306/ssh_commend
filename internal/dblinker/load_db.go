package dblinker

import (
	"ssh_commend/internal/dblinker/dbmd"
	"ssh_commend/internal/sysenv"
	"strconv"

	"bitbucket.org/okestrolab/baton-ao-sdk/btocd"

	"gorm.io/gorm"
)

func LoadModule(db *gorm.DB) ([]*sysenv.VmEnv, error) {

	var vmList []*dbmd.BtResource
	res := db.Raw("select * from bt_resource br where br.resource_type_id = ? and (br.del_yn = 'n' or br.del_yn = 'N')", btocd.ResourceType.OsVm).Scan(&vmList)
	if res.Error != nil {
		return nil, res.Error
	}

	VmsList := make([]*sysenv.VmEnv, len(vmList))

	for i, v := range vmList {
		VmsList[i] = &sysenv.VmEnv{
			ID:    v.ID,
			PrvId: v.ProviderID,
			UUID:  v.UuID,
			Name:  v.Name,
			Cfg:   new(sysenv.Configuration),
		}
	}

	return VmsList, nil
}

// main_LoadDb는 데이터베이스에서 설정을 로드하는 함수입니다.
func LoadDb(db *gorm.DB, c *sysenv.VmEnv) (err error) {

	// module attr
	{
		var dbList []*dbmd.BtResourceAttr
		res := db.Raw("select * from bt_resource_attr bra left join bt_resource br on bra.resource_id = br.id where (br.del_yn= 'n' or br.del_yn= 'N') and br.id = ? and br.resource_type_id = ? order by br.id asc", c.ID, btocd.ResourceType.OsVm).Scan(&dbList)
		if res.Error != nil {
			return res.Error
		}

		for _, v := range dbList {
			switch v.Key {
			case "vm_manage_ip":
				c.Cfg.Vmattr.HostIp = v.Value
			case "vm_manage_port":
				port, _ := strconv.ParseUint(v.Value, 10, 32)
				c.Cfg.Vmattr.Port = uint(port)
			case "vm_login_user":
				c.Cfg.Vmattr.Id = v.Value
			case "vm_login_pwd":
				c.Cfg.Vmattr.Pwd = v.Value
			case "OS-EXT-STS:power_state":
				c.Cfg.Vmattr.State, _= strconv.ParseFloat(v.Value, 64)
			}
		}
	}

	// module(logstash) attr
	{
		var dbList []*dbmd.BtModuleAttr
		res := db.Raw("select bma.* from bt_module bm inner join bt_module_attr bma on bm.id = bma.module_id where (bm.del_yn= 'n' or bm.del_yn= 'N') and bm.provider_id = ? and bm.module_type_id = ? order by bm.id asc", c.PrvId, btocd.ModuleType.LogstashFilter).Scan(&dbList)
		if res.Error != nil {
			return res.Error
		}

		for _, v := range dbList {
			switch v.Key {
			case "endpoints":
				c.Cfg.Logstash.Urls = v.Value
			}
		}
	}

	return nil
}

func UpdateDb(db *gorm.DB, c *sysenv.VmEnv, osinfo *sysenv.OsInfo) error {
	var vmList []*dbmd.BtResource

	res := db.Raw("select * from bt_resource br where br.resource_type_id = ? and br.`uuid` = ? and (br.del_yn = 'n' or br.del_yn = 'N')", btocd.ResourceType.OsVm, c.UUID).Scan(&vmList)

	if res.Error != nil {
		return res.Error
	}

	db.Clauses(clause.OnConflict{


}