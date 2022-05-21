package handler

import (
	"context"
	"micro-trainning-part4/cartOrder_srv/proto/pb"
	"micro-trainning-part4/custom_error"
	"micro-trainning-part4/internal"
	"micro-trainning-part4/shopcart_web/req"
	"net/http"
	"strconv"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func OrderListHandler(c *gin.Context) {
	accountIdStr := c.DefaultQuery("accountId", "0")
	accountId, _ := strconv.Atoi(accountIdStr)

	pageNoStr := c.DefaultQuery("pageNo", "0")
	pageNo, _ := strconv.Atoi(pageNoStr)

	pageSizeStr := c.DefaultQuery("pageSize", "0")
	pageSize, _ := strconv.Atoi(pageSizeStr)

	reqPb := &pb.OrderPagingReq{
		PageNo:    int32(pageNo),
		PageSize:  int32(pageSize),
		AccountId: int32(accountId),
	}

	// customClaims := claims.(*jwt_op.CustomClaims)
	// if customClaims.AuthorityId == 1 {
	// 	reqPb.AccountId = int32(accountId.(uint))
	// }

	ctx := context.WithValue(context.Background(), "webContext", c)

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
	idStr := c.DefaultQuery("id", "0")
	id, _ := strconv.Atoi(idStr)

	accountIdStr := c.DefaultQuery("accountId", "0")
	accountId, _ := strconv.Atoi(accountIdStr)
	res, err := internal.OrderClient.OrderDetail(context.WithValue(context.Background(), "webContext", c), &pb.OrderItemReq{
		Id:        int32(id),
		AccountId: int32(accountId),
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
	accountIdStr := c.DefaultQuery("accountId", "0")
	accountId, _ := strconv.Atoi(accountIdStr)
	rsp, err := internal.OrderClient.CreateOrder(context.WithValue(context.Background(), "webContext", c), &pb.OrderItemReq{
		AccountId: int32(accountId),
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
