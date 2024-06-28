package dbmd

import "time"

type VwResourceDetail struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement:true"`
	ProviderID     int64     `gorm:"column:provider_id;not null"`
	ResourceTypeID int64     `gorm:"column:resource_type_id;not null"`
	UuID           string    `gorm:"column:uuid;not null"`
	Name           string    `gorm:"column:name"`
	Description    string    `gorm:"column:description"`
	DelYn          string    `gorm:"column:del_yn;not null;default:N"`
	RegDt          time.Time `gorm:"column:reg_dt;default:now()"`
	ModDt          time.Time `gorm:"column:mod_dt"`
	DelDt          time.Time `gorm:"column:del_dt"`

	ResourceID int64     `gorm:"column:resource_id;not null"`
	Status     string    `gorm:"column:status;not null"`
	CreateDt   time.Time `gorm:"column:create_dt"`
	UpdateDt   time.Time `gorm:"column:update_dt"`
	DeleteDt   time.Time `gorm:"column:delete_dt"`
	UpdateVer  string    `gorm:"column:update_ver"`

	IVal struct {
		IsExist bool // fase 일 경우, db 에 delete 된다.
	} `gorm:"-"`
}

func (*VwResourceDetail) TableName() string {
	return "vw_resource_detail"
}
