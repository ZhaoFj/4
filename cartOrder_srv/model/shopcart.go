package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int32          `gorm:"primary_key"`
	CreateAt  *time.Time     `gorm:"autoCreateTime;column:add_time"`
	UpdateAt  *time.Time     `gorm:"autoUpdateTime;column:update_time"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ShopCart struct {
	BaseModel
	AccountId int32 `gorm:"type:int;index"`
	ProductId int32 `gorm:"type:int;index"`
	Num       int32 `gorm:"int"`
	Checked   bool
}
