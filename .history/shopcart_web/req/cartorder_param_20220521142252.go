package req

type ShopCartReq struct {
	Id        int32 `json:"id"`
	AccountId int32 `json:"account_id" binding:"required"`
	ProductId int32 `json:"product_id" binding:"required"`
	Num       int32 `json:"num" binding:"required"`
	Checked   bool  `json:"checked"`
}

type DelShopCartReq struct {
	Id        int32 `json:"id"`
	AccountId int32 `json:"account_id" binding:"required"`
	ProductId int32 `json:"product_id" binding:"required"`
}

type OrderReq struct {
	AccountId int32  `json:"accountId" binding:"required"`
	Receiver  string `json:"receiver" binding:"required"`
	Mobile    string `json:"mobile" binding:"required"`
	Addr      string `json:"addr" binding:"required"`
	PostCode  string `json:"postcode" binding:"required"`
}
