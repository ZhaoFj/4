package handler

import (
	"context"
	"net/http"
	"strconv"

	"micro-trainning-part4/cartOrder_srv/biz"
	"micro-trainning-part4/cartOrder_srv/proto/pb"
	"micro-trainning-part4/custom_error"
	"micro-trainning-part4/internal"
	"micro-trainning-part4/jwt_op"
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
	accountId, _ := c.Get("accountId")
	claims, _ := c.Get("claims")

	pageNoStr := c.DefaultQuery("pageNo", "0")
	pageNo, _ := strconv.Atoi(pageNoStr)

	pageSizeStr := c.DefaultQuery("pageSize", "0")
	pageSize, _ := strconv.Atoi(pageSizeStr)

	reqPb := &pb.OrderPagingReq{
		PageNo:   int32(pageNo),
		PageSize: int32(pageSize),
	}

	customClaims := claims.(*jwt_op.CustomClaims)
	if customClaims.AuthorityId == 1 {
		reqPb.AccountId = int32(accountId.(uint))
	}

	ctx := context.WithValue(context.Background(), "webContent", c)

	res, err := internal.OrderClient.OrderList(ctx, reqPb)
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
	id, _ := c.Get("id")
	accountId, _ := c.Get("accountId")
	res, err := internal.OrderClient.OrderDetail(context.WithValue(context.Background(), "webContent", c), &pb.OrderItemReq{
		Id:        int32(id.(uint)),
		AccountId: int32(accountId.(uint)),
	})
	if err != nil {
		zap.S().Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "获取订单失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":   "",
		"order": res.Order,
		"list":  res.ProductList,
	})
}

func CreaterOrderHandler(c *gin.Context) {
	orderReq := req.OrderReq{}
	if err := c.ShouldBindJSON(&orderReq); err != nil {
		zap.S().Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": custom_error.InternalServerError,
		})
		return
	}
	accountId, _ := c.Get("accountId")
	rsp, err := internal.OrderClient.CreateOrder(context.WithValue(context.Background(), "webContent", c), &pb.OrderItemReq{
		AccountId: int32(accountId.(uint)),
		Receiver:  orderReq.Receiver,
		Mobile:    orderReq.Mobile,
		Addr:      orderReq.Addr,
		PostCode:  orderReq.PostCode,
	})
	if err != nil {
		zap.S().Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "订单创建失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "",
		"id":  rsp.Id,
	})
}
