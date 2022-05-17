package internal

import (
	"encoding/json"
	"fmt"
	"micro-trainning-part4/cartOrder_srv/proto/pb"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	AppConf        AppConfig
	NacosConf      NacosConfig
	ShopCartClient pb.ShopCartServiceClient
	OrderClient    pb.OrderServiceClient
	ProductClient  pb.ProductServiceClient
	StockClient    pb.StockServiceClient
)

func init() {
	initNacos()
	initFromNacos()
	//fmt.Println("Nacos初始化完成")

}

func initNacos() {
	v := viper.New()
	//fmt.Println("Viper:", v)
	//fmt.Println(fileName)
	//v.SetConfigFile("config")
	v.SetConfigName("pro-config")
	v.AddConfigPath("$GOPATH/src/micro-trainning-part4/")
	v.SetConfigType("yml")
	err := v.ReadInConfig()

	if err != nil {
		panic(err)
	}
	v.Unmarshal(&NacosConf)
	//fmt.Println("NacosConf:", NacosConf)
}

func initFromNacos() {
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: NacosConf.Host,
			Port:   NacosConf.Port,
		},
	}

	LogRollingConfig := constant.ClientLogRollingConfig{
		MaxSize: 10,
		MaxAge:  3,
	}

	clientConfig := constant.ClientConfig{
		//NamespaceId:         "edfa58f5-77cf-49f1-add4-3459c8ccfe98",
		NamespaceId:         NacosConf.NameSpace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "nacos/log",
		CacheDir:            "nacos/cache",
		LogLevel:            "debug",
		LogRollingConfig:    &LogRollingConfig,
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: NacosConf.DataId,
		Group:  NacosConf.Group,
	})
	if err != nil {
		panic(err)
	}
	//fmt.Println(content)

	json.Unmarshal([]byte(content), &AppConf)
	//fmt.Println(AppConf)
}

func initGrpcClient() {
	addr := fmt.Sprintf("%s:%d", AppConf.ConsulConfig.Host, AppConf.ConsulConfig.Port)
	dailAddr := fmt.Sprintf("consul://%s/%s?wait=14s", addr, AppConf.ShopCartSrvConfig.SrvName)
	conn, err := grpc.Dial(
		dailAddr,
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		s := fmt.Sprintf("ShopCart-GRPC拨号失败:%s", err.Error())
		zap.S().Fatal(s)
		panic(err)
	}

	ShopCartClient = pb.NewShopCartServiceClient(conn)
	OrderClient = pb.NewOrderServiceClient(conn)

	productSrvAddr := fmt.Sprintf("%s:%d", AppConf.ProductSrvConfig.Host, AppConf.ProductSrvConfig.Port)
	productConn, err = grpc.Dial(
		productSrvAddr,
		grpc.WithInsecure(),
	)

	if err != nil {
		s := fmt.Sprintf("productSrv-GRPC拨号失败:%s", err.Error())
		zap.S().Fatal(s)
		panic(err)
	}
	ProductClient = pb.NewProductServiceClient(productConn)

}
