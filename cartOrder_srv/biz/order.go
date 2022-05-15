package biz

import (
	"context"
	"errors"
	"micro-trainning-part4/cartOrder_srv/model"
	"micro-trainning-part4/cartOrder_srv/proto/pb"
	"micro-trainning-part4/custom_error"
	"micro-trainning-part4/internal"
	"micro-trainning-part4/util"

	"google.golang.org/protobuf/types/known/emptypb"
)

//新建订单
func (s CartOrderServer) CreateOrder(c context.Context, req *pb.OrderItemReq) (*pb.OrderItemRes, error) {
	panic("xxx")
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
