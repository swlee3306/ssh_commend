package dbmd

import "time"

type BtResourceStat struct {
	ResourceID int64     `gorm:"column:resource_id;primaryKey"`
	Status     string    `gorm:"column:status;not null"`
	CreateDt   time.Time `gorm:"column:create_dt"`
	UpdateDt   time.Time `gorm:"column:update_dt"`
	DeleteDt   time.Time `gorm:"column:delete_dt"`
	UpdateVer  string    `gorm:"column:update_ver"`
}

func (*BtResourceStat) TableName() string {
	return "bt_resource_stat"
}
