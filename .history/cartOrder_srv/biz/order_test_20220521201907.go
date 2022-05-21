package biz

import (
	"fmt"
	"micro-trainning-part4/cartOrder_srv/proto/pb"
	"micro-trainning-part4/internal"

	"google.golang.org/grpc"
)

var orderServiceClient pb.OrderServiceClient

func init() {
	addr := fmt.Sprintf("%s:%d", internal.AppConf.ShopCartSrvConfig.Host, internal.AppConf.ShopCartSrvConfig.Port)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic("grpc初始化失败")
	}
	orderServiceClient = pb.NewOrderServiceClient(conn)
}

func TestOrderServer_OrderList(t *Testing.t) {

}
