package biz

import (
	"context"
	"errors"
	"micro-trainning-part4/cartOrder_srv/model"
	"micro-trainning-part4/cartOrder_srv/proto/pb"
	"micro-trainning-part4/custom_error"
	"micro-trainning-part4/internal"
	"micro-trainning-part4/util"

	"github.com/nacos-group/nacos-sdk-go/inner/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

//新建订单
func (s CartOrderServer) CreateOrder(c context.Context, req *pb.OrderItemReq) (*pb.OrderItemRes, error) {
	/*
		1 拿到购物车内选定的商品
		2 计算订单总金额
		3 扣减库存
		4 将订单数据写入数据库 orderItem 和 orderProduct 表
		5 删除购物车内已购买商品
	*/
	var productIds []int32
	var cartList []model.ShopCart
	productNumMap := make(map[int32]int32)
	r := internal.DB.Where(&model.ShopCart{AccountId: req.AccountId, Checked: true}).Find(&cartList)
	if r.RowsAffected == 0 {
		return nil, errors.New(custom_error.ProductNotChecked)
	}
	for _, item := range cartList {
		productIds = append(productIds, item.ProductId)
		productNumMap[item.ProductId] = item.Num
	}
	// res, err := internal.ProductClient.BatchGetProduct(context.Background(), &pb.BatchProductIdReq{Ids: productIds})
	// if err != nil {
	// 	zap.S().Error("[BatchGetProduct调用失败]", err)
	// 	return nil, errors.New(custom_error.InternalServerError)
	// }
	res := &pb.ProductRes{
		Total: 1,
		ItemList: []*pb.ProductItemRes{
			{
				Id:         3,
				Name:       "商品2",
				CoverImage: "cover_img_url",
				RealPrice:  22.2,
			},
		},
	}

	//计算价格
	var amount float32
	var stockItemList []*pb.ProductStockItem
	var orderProductList []model.OrderProduct
	for _, item := range res.ItemList {
		amount += item.RealPrice * float32(productNumMap[item.Id])
		stockItemList = append(stockItemList, &pb.ProductStockItem{ProductId: item.Id, Num: productNumMap[item.Id]})
		var orderProduct = model.OrderProduct{
			OrderId:     req.AccountId,
			ProductId:   item.Id,
			ProductName: item.Name,
			CoverImg:    item.CoverImage,
			RealPrice:   item.RealPrice,
			Num:         productNumMap[item.Id],
		}
		orderProductList = append(orderProductList, orderProduct)
	}

	//扣减库存
	_, err := internal.StockClient.Sell(context.Background(), &pb.SellItem{StockItemList: stockItemList})
	if err != nil {
		return nil, errors.New(custom_error.StockNotEnough)
	}

	//创建订单
	tx := internal.DB.Begin()
	var orderItem model.OrderItem
	orderItem.AccountId = req.Id
	uuid, _ := uuid.NewV4()
	orderItem.OrderNum = uuid.String()
	orderItem.Status = "unPay"
	orderItem.Addr = req.Addr
	orderItem.Receiver = req.Receiver
	orderItem.ReceiverMobile = req.Mobile
	orderItem.PostCode = req.PostCode
	orderItem.OrderAmount = amount

	result := tx.Save(&orderItem)
	if result.Error != nil || result.RowsAffected < 1 {
		tx.Rollback()
		return nil, errors.New(custom_error.CreateOrderFailed)
	}

	for _, item := range orderProductList {
		item.OrderId = orderItem.ID
	}
	result = tx.CreateInBatches(orderProductList, 50)
	if result.Error != nil || result.RowsAffected < 1 {
		tx.Rollback()
		return nil, errors.New(custom_error.CreateOrderFailed)
	}

	//删除购物车已下单商品
	result = tx.Where(&model.ShopCart{Checked: true, AccountId: req.AccountId}).Delete(&model.ShopCart{})
	if result.Error != nil || result.RowsAffected < 1 {
		tx.Rollback()
		return nil, errors.New(custom_error.CreateOrderFailed)
	}
	tx.Commit()

	orderItemRes := ConverOrderItemModel2Pb(orderItem)

	return orderItemRes, nil
}

//订单列表
func (s CartOrderServer) OrderList(c context.Context, req *pb.OrderPagingReq) (*pb.OrderListRes, error) {
	var orderList []model.OrderItem
	var res pb.OrderListRes
	var total int64

	internal.DB.Where(&model.OrderItem{AccountId: req.AccountId}).Count(&total)
	res.Total = int32(total)

	internal.DB.Where(
		&model.OrderItem{
			AccountId: req.AccountId,
		},
	).Scopes(util.Paginate(int(req.PageNo), int(req.PageSize))).Find(&orderList)
	for _, item := range orderList {
		r := ConverOrderItemModel2Pb(item)
		res.Itemlist = append(res.Itemlist, r)
	}

	return &res, nil
}

//订单详情
func (s CartOrderServer) OrderDetail(c context.Context, req *pb.OrderItemReq) (*pb.OrderItemDetailRes, error) {
	var orderDetail model.OrderItem
	var detailRes pb.OrderItemDetailRes
	r := internal.DB.Where(&model.OrderItem{BaseModel: model.BaseModel{ID: req.Id}, AccountId: req.AccountId}).First(&orderDetail)
	if r.RowsAffected == 0 {
		return nil, errors.New(custom_error.OrderNotFound)
	}
	res := ConverOrderItemModel2Pb(orderDetail)
	detailRes.Order = res
	var orderProductList []model.OrderProduct
	internal.DB.Where(&model.OrderProduct{OrderId: orderDetail.ID}).Find(&orderProductList)
	for _, product := range orderProductList {
		orderProudctRes := ConverOrderProductModel2Pb(product)
		detailRes.ProductList = append(detailRes.ProductList, orderProudctRes)
	}
	return &detailRes, nil
}

//更改状态
func (s CartOrderServer) ChangeOrderStatus(c context.Context, req *pb.OrderStatus) (*emptypb.Empty, error) {
	r := internal.DB.Model(&model.OrderItem{}).Where("order_no=?", req.OrderNum).Update("status=?", req.Status)
	if r.RowsAffected == 0 {
		return nil, errors.New(custom_error.OrderNotFound)
	}
	return &emptypb.Empty{}, nil
}

func ConverOrderItemModel2Pb(req model.OrderItem) *pb.OrderItemRes {
	var res pb.OrderItemRes
	res.Id = req.ID
	res.AccountId = req.AccountId
	res.Addr = req.Addr
	res.Amount = req.OrderAmount
	res.CreateTime = req.CreateAt.String()
	res.Mobile = req.ReceiverMobile
	res.OrderNo = req.OrderNum
	res.PayType = req.PayType
	res.PostCode = req.PostCode
	res.Receiver = req.Receiver
	res.Status = req.Status
	return &res
}

func ConverOrderProductModel2Pb(req model.OrderProduct) *pb.OrderProductRes {
	var res pb.OrderProductRes
	res.Id = req.ID
	res.CoverImg = req.CoverImg
	res.Num = req.Num
	res.OrderId = req.OrderId
	res.ProductId = req.ProductId
	res.ProductName = req.ProductName
	res.RealPrice = req.RealPrice
	return &res
}
