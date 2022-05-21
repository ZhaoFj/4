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
