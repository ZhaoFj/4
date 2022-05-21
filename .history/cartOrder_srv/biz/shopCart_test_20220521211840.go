package biz

import (
	"context"
	"fmt"
	"micro-trainning-part4/cartOrder_srv/proto/pb"
	"micro-trainning-part4/internal"
	"testing"
)

func TestShopCartServer_AddShopCartItem(t *testing.T) {
	shopCart := pb.ShopCartReq{
		ProductId: 1,
		AccountId: 1,
		Num:       1,
		Checked:   false,
	}
	res, err := internal.ShopCartClient.AddShopCartItem(context.Background(), &shopCart)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)

	for i := 1; i < 6; i++ {
		shopCart1 := pb.ShopCartReq{
			ProductId: 2,
			AccountId: 2,
			Num:       1,
			Checked:   false,
		}
		res, err := internal.ShopCartClient.AddShopCartItem(context.Background(), &shopCart1)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(res)
	}
}

func TestShopCartServer_UpdateShopCartItem(t *testing.T) {
	shopCart := pb.ShopCartReq{
		ProductId: 2,
		AccountId: 2,
		Num:       1,
		Checked:   true,
	}
	res, err := internal.ShopCartClient.UpdateShopCartItem(context.Background(), &shopCart)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestShopCartServer_ShopCartItemList(t *testing.T) {
	req := pb.AccountReq{
		AccountId: 1,
	}
	res, err := internal.ShopCartClient.ShopCartItemList(context.Background(), &req)
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range res.ItemList {
		fmt.Println(item.ProductId, ": ", item.Num)
	}
}
