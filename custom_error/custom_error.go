package custom_error

const (
	ParamError             = "参数错误"
	InternalServerError    = "服务器内部错误"
	ProductNotFound        = "产品不存在"
	StockNotEnough         = "产品余量不足"
	AddShopCartItemFail    = "购物车添加失败"
	NotInShopCart          = "购物车中无此商品"
	UpdateShopCartItemFail = "购物车更新失败"
	DeleteShopCartItemFail = "购物车删除失败"
	GetShopCartListFail    = "购物车列表获取失败"
	OrderNotFound          = "查无此单"
)
