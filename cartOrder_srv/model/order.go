package model

import "time"

type OrderItem struct {
	BaseModel
	AccountId      int32  `gorm:"type:int;index"`
	OrderNum       string `gorm:"type:varchar(64);index"` //订单号
	PayType        string `gorm:"type:varchar(16)"`       //支付方式
	Status         string `gorm:"type:varchar(16)"`       //支付状态
	TradeNo        string `gorm:"type:varchar(64)"`       //支付流水
	Addr           string `gorm:"type:varchar(64)"`
	Receiver       string `gorm:"type:varchar(16)"`
	ReceiverMobile string `gorm:"type:varchar(11)"`
	PostCode       string `gorm:"type:varchar(16)"`
	OrderAmount    float32
	PayTime        *time.Time `gorm:"type:datetime"`
}

type OrderProduct struct {
	BaseModel
	OrderId     int32  `gorm:"type:int;index"`
	ProductId   int32  `gorm:"type:int;index"`
	ProductName string `gorm:"type:varchar(64);index"`
	CoverImg    string `gorm:"type:varchar(128)"`
	RealPrice   float32
	Num         int32 `gorm:"type:int"`
}
