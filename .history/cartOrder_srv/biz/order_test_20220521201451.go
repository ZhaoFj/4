package biz

import "micro-trainning-part4/cartOrder_srv/proto/pb"

var orderServiceClient pb.OrderServiceClient

func init() {
	addr := fmt.Sprintf("%s:%d", internal.AppConf, internal.AppConf.ShopCartSrvConfig.Port)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic("grpc初始化失败")
	}
	orderServiceClient = pb.NewOrderServiceClient(conn)
}

func 