package handler

import (
	"context"
	"net/http"
	"strconv"

	"micro-trainning-part4/cartOrder_srv/biz"
	"micro-trainning-part4/cartOrder_srv/proto/pb"
	"micro-trainning-part4/custom_error"
	"micro-trainning-part4/internal"
	"micro-trainning-part4/shopcart_web/req"

	"github.com/gin-gonic/gin"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
)

func CartOrderListHandler(c *gin.Context) {
	accountIdStr := c.DefaultQuery("accountId", "0")
	accountId, err := strconv.Atoi(accountIdStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": custom_error.ParamError,
		})
		return
	}
	var req pb.AccountReq
	req.AccountId = int32(accountId)
	res, err := internal.ShopCartClient.ShopCartItemList(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": custom_error.GetShopCartListFail,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":   "",
			"total": res.Total,
			"list":  res.ItemList,
		})
	}
}

func AddShopCartItemHandler(c *gin.Context) {
	var shopCartReq req.ShopCartReq
	err := c.ShouldBindJSON(&shopCartReq)
	if err != nil {
		zap.S().Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": custom_error.ParamError,
		})
		return
	}
	r := biz.ConverShopCartReq2pb(shopCartReq)
	res, err := internal.ShopCartClient.AddShopCartItem(context.Background(), r)
	if err != nil {
		zap.S().Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": custom_error.AddShopCartItemFail,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "",
		"res": res,
	})
}

func UpdateShopCartItemHandler(c *gin.Context) {
	var shopCartReq req.ShopCartReq
	err := c.ShouldBindJSON(&shopCartReq)
	if err != nil {
		zap.S().Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": custom_error.ParamError,
		})
		return
	}
	r := biz.ConverShopCartReq2pb(shopCartReq)
	_, err = internal.ShopCartClient.UpdateShopCartItem(context.Background(), r)
	if err != nil {
		zap.S().Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": custom_error.UpdateShopCartItemFail,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "",
	})
}

func DeleteShopCartItemHandler(c *gin.Context) {
	var delShopCartReq req.DelShopCartReq
	err := c.ShouldBindJSON(&delShopCartReq)
	if err != nil {
		zap.S().Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": custom_error.ParamError,
		})
		return
	}
	r := biz.ConverDelShopCartReq2pb(delShopCartReq)
	_, err = internal.ShopCartClient.DeleteShopCartItem(context.Background(), r)
	if err != nil {
		zap.S().Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": custom_error.DeleteShopCartItemFail,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "",
	})
}

func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func OrderListHandler(c *gin.Context) {
	res, err := internal.OrderClient.OrderList(context.Background(), &pb.OrderPagingReq{
		AccountId: 0,
		PageNo:    0,
		PageSize:  0,
	})
	if err != nil {
		zap.S().Error("OrderList错误:", err)
		c.JSON(http.StatusOK, gin.H{
			"msg": custom_error.InternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":   "",
		"items": res.Itemlist,
	})
}

func OrderDetailHandler(c *gin.Context) {

}

func CreaterOrderHandler(c *gin.Context) {

}
