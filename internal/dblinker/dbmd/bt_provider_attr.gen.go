// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dbmd

import (
	"time"
)

const TableNameBtProviderAttr = "bt_provider_attr"

// BtProviderAttr mapped from table <bt_provider_attr>
type BtProviderAttr struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ProviderID  int64     `gorm:"column:provider_id;not null" json:"provider_id"`
	Key         string    `gorm:"column:key;not null" json:"key"`
	Value       string    `gorm:"column:value" json:"value"`
	Description string    `gorm:"column:description" json:"description"`
	DelYn       string    `gorm:"column:del_yn;default:N" json:"del_yn"`
	RegDt       time.Time `gorm:"column:reg_dt;default:current_timestamp()" json:"reg_dt"`
	ModDt       time.Time `gorm:"column:mod_dt" json:"mod_dt"`
	DelDt       time.Time `gorm:"column:del_dt" json:"del_dt"`
}

// TableName BtProviderAttr's table name
func (*BtProviderAttr) TableName() string {
	return TableNameBtProviderAttr
}