// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dbmd

import (
	"time"
)

const TableNameBtCollectorType = "bt_collector_type"

// BtCollectorType mapped from table <bt_collector_type>
type BtCollectorType struct {
	ID          int64     `gorm:"column:id;primaryKey" json:"id"`
	Name        string    `gorm:"column:name;not null" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	DelYn       string    `gorm:"column:del_yn;default:N" json:"del_yn"`
	RegDt       time.Time `gorm:"column:reg_dt;default:current_timestamp()" json:"reg_dt"`
	ModDt       time.Time `gorm:"column:mod_dt" json:"mod_dt"`
	DelDt       time.Time `gorm:"column:del_dt" json:"del_dt"`
}

// TableName BtCollectorType's table name
func (*BtCollectorType) TableName() string {
	return TableNameBtCollectorType
}
