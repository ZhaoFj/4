package biz

import (
	"context"
	"errors"
	"micro-trainning-part4/cartOrder_srv/model"
	"micro-trainning-part4/cartOrder_srv/proto/pb"
	"micro-trainning-part4/custom_error"
	"micro-trainning-part4/internal"
	"micro-trainning-part4/shopcart_web/req"

	"google.golang.org/protobuf/types/known/emptypb"
)

type CartOrderServer struct {
}

//购物车列表
func (s CartOrderServer) ShopCartItemList(c context.Context, req *pb.AccountReq) (*pb.CartItemListRes, error) {
	var cartItemList []model.ShopCart
	var res pb.CartItemListRes
	var itemList []*pb.CartItemRes
	r := internal.DB.Where(&model.ShopCart{AccountId: req.AccountId}).Find(&cartItemList)
	if r.Error != nil {
		return nil, errors.New(custom_error.ParamError)
	}
	if r.RowsAffected < 1 {
		return &res, nil
	}

	for _, item := range cartItemList {
		itemPb := ConvertShopCartModel2Pb(item)
		itemList = append(itemList, itemPb)
	}
	res.Total = int32(r.RowsAffected)
	res.ItemList = itemList
	return &res, nil
}

//添加产品到购物车
func (s CartOrderServer) AddShopCartItem(c context.Context, req *pb.ShopCartReq) (*pb.CartItemRes, error) {
	var shopCart model.ShopCart
	var res *pb.CartItemRes
	r := internal.DB.Where(&model.ShopCart{AccountId: req.AccountId, ProductId: req.ProductId}).First(&shopCart)
	if r.RowsAffected == 1 {
		shopCart.Num += req.Num
	}
	if r.RowsAffected < 1 {
		shopCart.AccountId = req.AccountId
		shopCart.ProductId = req.ProductId
		shopCart.Num = req.Num
		shopCart.Checked = req.Checked
	}
	r = internal.DB.Save(&shopCart)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.AddShopCartItemFail)
	}
	res = ConvertShopCartModel2Pb(shopCart)
	return res, nil
}

//更新购物车内的产品
func (s CartOrderServer) UpdateShopCartItem(c context.Context, req *pb.ShopCartReq) (*emptypb.Empty, error) {
	var shopCart model.ShopCart
	r := internal.DB.Where("account_id=? and product_id=?", req.AccountId, req.ProductId).First(&shopCart)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.NotInShopCart)
	}
	if req.Num < 1 {
		return nil, errors.New(custom_error.ParamError)
	}
	shopCart.Num = req.Num
	shopCart.Checked = req.Checked
	r = internal.DB.Save(&shopCart)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.AddShopCartItemFail)
	}
	return &emptypb.Empty{}, nil
}

//删除购物车内的产品
func (s CartOrderServer) DeleteShopCartItem(c context.Context, req *pb.DelShopCartReq) (*emptypb.Empty, error) {
	var shopCart model.ShopCart
	r := internal.DB.Where("account_id=? and product_id=?", req.AccountId, req.ProductId).First(&shopCart)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.NotInShopCart)
	}
	r = internal.DB.Delete(&shopCart)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.DeleteShopCartItemFail)
	}
	return &emptypb.Empty{}, nil
}

func ConvertShopCartModel2Pb(req model.ShopCart) *pb.CartItemRes {
	var res pb.CartItemRes
	res.AccountId = req.AccountId
	res.ProductId = req.ProductId
	res.Id = req.ID
	res.Num = req.Num
	res.Checked = req.Checked

	return &res
}

func ConverShopCartReq2pb(req req.ShopCartReq) *pb.ShopCartReq {
	var res pb.ShopCartReq
	res.Id = req.Id
	res.AccountId = req.AccountId
	res.ProductId = req.ProductId
	res.Num = req.Num
	res.Checked = req.Checked
	return &res
}

func ConverDelShopCartReq2pb(req req.DelShopCartReq) *pb.DelShopCartReq {
	var res pb.DelShopCartReq
	res.Id = req.Id
	res.AccountId = req.AccountId
	res.ProductId = req.ProductId
	return &res
}
