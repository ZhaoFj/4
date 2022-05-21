package biz

import (
	"context"
	"fmt"
	"micro-trainning-part4/cartOrder_srv/proto/pb"
	"micro-trainning-part4/internal"
	"testing"
)

func TestShopCartServer_OrderList(t *testing.T) {
	res, err := internal.OrderClient.OrderList(context.Background(), &pb.OrderPagingReq{
		AccountId: 1,
		PageNo:    1,
		PageSize:  5,
	})
	if err != nil {
		panic(err)
	}
	for _, item := range res.Itemlist {
		fmt.Println(item.Id, "_", item.AccountId)
	}
}

func TestShopCartServer_CreateOrder(t *testing.T) {
	res, err := internal.OrderClient.CreateOrder(context.Background(), &pb.OrderItemReq{
		AccountId: 1,
		Addr:      "翻斗花园",
		PostCode:  "000000",
		Receiver:  "俺滴图图",
		Mobile:    "123456",
		PayType:   "未支付",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Id)
}
