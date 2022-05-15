package internal

import (
	"encoding/json"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
)

var AppConf AppConfig
var NacosConf NacosConfig

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

func init() {
	initNacos()
	initFromNacos()
	//fmt.Println("Nacos初始化完成")

}
