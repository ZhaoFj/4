package biz

import (
	"context"
	"fmt"
	"micro-trainning-part4/cartOrder_srv/proto/pb"
	"micro-trainning-part4/internal"
	"testing"
)

func TestOrderServer_OrderList(t *testing.T) {
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
