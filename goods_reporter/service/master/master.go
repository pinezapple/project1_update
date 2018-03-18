package master

import (
	"log"
	rpc "project1_update/goods_reporter/Goodreportrpc"
	goodsDB "project1_update/goods_reporter/microservice/goodsDB"

	"golang.org/x/net/context"
)

type MasterServerImp struct{}

func (c *MasterServerImp) IfGoodExist(ctx context.Context, req *rpc.IfGoodExistReq) (*rpc.IfGoodExistResp, error) {
	r, err := goodsDB.IfGoodExist(req.Id)
	if err != nil {
		log.Fatalf("can't check whether goods exist in master server of goods_reporter %v", err)
		return &rpc.IfGoodExistResp{Ok: false}, err
	}
	return &rpc.IfGoodExistResp{Ok: r}, err
}

func (c *MasterServerImp) SelectAllGood(ctx context.Context, req *rpc.SelectAllGoodReq) (*rpc.SelectAllGoodResp, error) {
	r, err := goodsDB.SelectAllGoods()
	if err != nil {
		log.Fatalf("can't select all from goods_reporter %v", err)
		return &rpc.SelectAllGoodResp{Payload: nil}, err
	}
	return &rpc.SelectAllGoodResp{Payload: r}, err
}

func (c *MasterServerImp) SelectGoodById(ctx context.Context, req *rpc.SelectGoodByIdReq) (*rpc.SelectGoodByIdResp, error) {
	r, err := goodsDB.SelectGoodsById(req.Id)
	if err != nil {
		log.Fatalf("can't cus by id from goods_reporter %v", err)
		return &rpc.SelectGoodByIdResp{Payload: nil}, err
	}
	return &rpc.SelectGoodByIdResp{Payload: r}, err
}

func (c *MasterServerImp) SelectGoodByName(ctx context.Context, req *rpc.SelectGoodByNameReq) (*rpc.SelectGoodByNameResp, error) {
	r, err := goodsDB.SelectGoodsByName(req.Name)
	if err != nil {
		log.Fatalf("can't cus by name from goods_reporter %v", err)
		return &rpc.SelectGoodByNameResp{Payload: nil}, err
	}
	return &rpc.SelectGoodByNameResp{Payload: r}, err
}

func (c *MasterServerImp) SelectGoodByPrice(ctx context.Context, req *rpc.SelectGoodByPriceReq) (*rpc.SelectGoodByPriceResp, error) {
	r, err := goodsDB.SelectGoodsByPrice(req.Price)
	if err != nil {
		log.Fatalf("can't cus by price from goods_reporter %v", err)
		return &rpc.SelectGoodByPriceResp{Payload: nil}, err
	}
	return &rpc.SelectGoodByPriceResp{Payload: r}, err
}
