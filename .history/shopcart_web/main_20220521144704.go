package main

import (
	"flag"
	"fmt"
	"micro-trainning-part4/internal"
	"micro-trainning-part4/internal/register"
	"micro-trainning-part4/shopcart_web/handler"
	"micro-trainning-part4/shopcart_web/middleware"
	"micro-trainning-part4/util"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	consulRegistry register.ConsulRegistry
	randomId       string
)

func init() {
	randomPort := util.GenRandomPort()
	if !internal.AppConf.Debug {
		internal.AppConf.ShopCartWebConfig.Port = randomPort
	}
	randomId = uuid.New().String()
	consulRegistry = register.NewConsulRegistry(internal.AppConf.ShopCartWebConfig.Host,
		internal.AppConf.ShopCartWebConfig.Port)

	consulRegistry.Register(internal.AppConf.ShopCartWebConfig.SrvName,
		randomId,
		internal.AppConf.ShopCartWebConfig.Host,
		internal.AppConf.ShopCartWebConfig.Port,
		internal.AppConf.ShopCartWebConfig.Tags)
}

func main() {
	ip := flag.String("ip", "0.0.0.0", "输入IP")
	//ip := internal.AppConf.ProductWebConfig.Host
	port := internal.AppConf.ShopCartWebConfig.Port
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *ip, port)
	r := gin.Default()
	cartGroup := r.Group("/v1/shopcart").Use(middleware.Tracing())
	{
		cartGroup.GET("/list", handler.CartOrderListHandler)
		cartGroup.POST("/add", handler.AddShopCartItemHandler)
		cartGroup.POST("/update", handler.UpdateShopCartItemHandler)
		cartGroup.POST("/delete", handler.DeleteShopCartItemHandler)
	}
	orderGroup := r.Group("/v1/order").Use(middleware.Tracing())
	{
		orderGroup.GET("/list", handler.OrderListHandler)
		orderGroup.GET("/:id", handler.OrderDetailHandler)
		orderGroup.POST("/add", handler.CreaterOrderHandler)
	}
	r.GET("/health", handler.HealthHandler)

	go func() {
		err := r.Run(addr)
		if err != nil {
			zap.S().Panic(addr + "启动失败" + err.Error())
		} else {
			zap.S().Info("启动成功:" + randomId)
		}
	}()
	q := make(chan os.Signal)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)
	<-q
	err := consulRegistry.DeRegister(randomId)
	if err != nil {
		zap.S().Panic("注销失败" + randomId + ":" + err.Error())
	} else {
		zap.S().Info("注销成功:" + randomId)
	}
}
