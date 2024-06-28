package dbmd

import "time"

type BtAgent struct {
	ID              int64  `gorm:"column:id;primaryKey;autoIncrement:true"`
	ResourceID      int64  `gorm:"column:resource_id;not null"`
	AgentKindID     int64  `gorm:"column:agent_kind_id;not null"`
	AgentProgramID  int64  `gorm:"column:agent_program_id"`
	AgentStatus     string `gorm:"column:agent_status"`
	AgentStatusPrev string `gorm:"column:agent_status_prev"`
	AgentStatusDesc string `gorm:"column:agent_status_desc"`

	UseYn string    `gorm:"column:use_yn;not null;default:Y"`
	DelYn string    `gorm:"column:del_yn;not null;default:N"`
	RegDt time.Time `gorm:"column:reg_dt;default:now()"`
	ModDt time.Time `gorm:"column:mod_dt"`
}

func (*BtAgent) TableName() string {
	return "bt_agent"
}
