package dbmd

import "time"

type BtResourceAttr struct {
	ID          int64     `gorm:"column:id;primaryKey"`
	ResourceID  int64     `gorm:"column:resource_id;not null"`
	Key         string    `gorm:"column:key;not null"`
	Value       string    `gorm:"column:value"`
	Description string    `gorm:"column:description"`
	DelYn       string    `gorm:"column:del_yn;not null"`
	RegDt       time.Time `gorm:"column:reg_dt"`
	ModDt       time.Time `gorm:"column:mod_dt"`
	DelDt       time.Time `gorm:"column:del_dt"`
}

func (*BtResourceAttr) TableName() string {
	return "bt_resource_attr"
}
