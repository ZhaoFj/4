package main

import (
	"fmt"
	"micro-trainning-part4/cartOrder_srv/biz"
	"micro-trainning-part4/cartOrder_srv/proto/pb"
	"micro-trainning-part4/internal"
	"net"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/hashicorp/consul/api"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func init() {
	//internal.InitViper("internal/dev-config.yml")
	internal.InitDB()
}

func main() {
	/*
		1.生成proto对应文件
		2.建立biz目录，实现接口
		3.拷贝之前的main文件中的函数
	*/

	port := internal.AppConf.ShopCartSrvConfig.Port
	addr := fmt.Sprintf("%s:%d", "0.0.0.0", port)

	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "0.0.0.0:6831",
		},
		ServiceName: "hantaMall",
	}
	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	defer closer.Close()
	if err != nil {
		panic(err)
	}

	opentracing.SetGlobalTracer(tracer)

	server := grpc.NewServer(grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(tracer)))
	pb.RegisterShopCartServiceServer(server, &biz.CartOrderServer{})
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		zap.S().Error("cartorder_srv启动异常:" + err.Error())
		panic(err)
	}
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	defaultConfig := api.DefaultConfig()
	defaultConfig.Address = fmt.Sprintf("%s:%d",
		internal.AppConf.ConsulConfig.Host,
		internal.AppConf.ConsulConfig.Port)
	client, err := api.NewClient(defaultConfig)
	if err != nil {
		panic(err)
	}
	checkAddr := fmt.Sprintf("%s:%d",
		internal.AppConf.ShopCartSrvConfig.Host,
		port)
	check := &api.AgentServiceCheck{
		GRPC:                           checkAddr,
		Timeout:                        "3s",
		Interval:                       "1s",
		DeregisterCriticalServiceAfter: "5s",
	}
	//随机一个UUID
	randUUID := uuid.New().String()
	fmt.Println(randUUID, "启动在", port)
	reg := api.AgentServiceRegistration{
		Name:    internal.AppConf.ShopCartSrvConfig.SrvName,
		ID:      randUUID,
		Port:    port,
		Tags:    internal.AppConf.ShopCartSrvConfig.Tags,
		Address: internal.AppConf.ShopCartSrvConfig.Host,
		Check:   check,
	}

	err = client.Agent().ServiceRegister(&reg)
	if err != nil {
		panic(err)
	}

	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
