package biz

import (
	"context"
	"micro-trainning-part4/cartOrder_srv/model"
	"micro-trainning-part4/cartOrder_srv/proto/pb"
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

	panic("xxx")
}

//更改状态
func (s CartOrderServer) ChangeOrderStatus(c context.Context, req *pb.OrderStatus) (*emptypb.Empty, error) {
	panic("xxx")
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
