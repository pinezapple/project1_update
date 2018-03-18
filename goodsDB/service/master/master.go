package master

import (
	"log"
	rpc "project1_update/goodsDB/GoodsDBrpc"
	"project1_update/goodsDB/dao"
	"project1_update/goodsDB/lib/MsgPack"

	goodsModel "project1_update/others/models/goodsmod"

	"golang.org/x/net/context"
)

type MasterServerImp struct{}

func (c *MasterServerImp) SelectAllGoods(ctx context.Context, req *rpc.SelectAllReq) (*rpc.SelectAllResp, error) {
	r, err := dao.GoodsDao.SelectAll()
	if err != nil {
		log.Fatalf("can't select all goods from goodsDB service %v", err)
		return &rpc.SelectAllResp{Payload: nil}, err
	}
	mar, _ := MsgPack.Marshal(r)
	return &rpc.SelectAllResp{Payload: mar}, nil
}

func (c *MasterServerImp) SelectByGoodsID(ctx context.Context, req *rpc.SelectByIDReq) (*rpc.SelectByIDResp, error) {
	r, ok, err := dao.GoodsDao.SelectByID(req.Id)
	if err != nil {
		log.Fatalf("can't select goodstomers by id from goodsDB service %v", err)
		return &rpc.SelectByIDResp{Payload: nil}, err
	}
	if ok == false {
		log.Fatalf("goods does not exist")
		return &rpc.SelectByIDResp{Payload: nil}, nil
	}
	mar, _ := MsgPack.Marshal(r)
	return &rpc.SelectByIDResp{Payload: mar}, nil
}

func (c *MasterServerImp) SelectByGoodsName(ctx context.Context, req *rpc.SelectByNameReq) (*rpc.SelectByNameResp, error) {
	r, ok, err := dao.GoodsDao.SelectByName(req.Name)
	if err != nil {
		log.Fatalf("can't select goods by name from goodsDB service %v", err)
		return &rpc.SelectByNameResp{Payload: nil}, err
	}
	if ok == false {
		log.Fatalf("goods does not exist")
		return &rpc.SelectByNameResp{Payload: nil}, nil
	}
	mar, _ := MsgPack.Marshal(r)
	return &rpc.SelectByNameResp{Payload: mar}, nil
}

func (c *MasterServerImp) SelectByGoodsPrice(ctx context.Context, req *rpc.SelectByGoodsPriceReq) (*rpc.SelectByGoodsPriceResp, error) {
	r, ok, err := dao.GoodsDao.SelectByPrice(req.Price)
	if err != nil {
		log.Fatalf("can't select goods by price from goodsDB service %v", err)
		return &rpc.SelectByGoodsPriceResp{Payload: nil}, err
	}
	if ok == false {
		log.Fatalf("goods does not exist")
		return &rpc.SelectByGoodsPriceResp{Payload: nil}, nil
	}
	mar, _ := MsgPack.Marshal(r)
	return &rpc.SelectByGoodsPriceResp{Payload: mar}, nil
}

func (c *MasterServerImp) AddGoods(ctx context.Context, req *rpc.AddReq) (*rpc.AddResp, error) {
	var obj *goodsModel.GoodsMod

	obj.Name = req.Name
	obj.Quantity = req.Quantity
	obj.Price = req.Price

	ok, err := dao.GoodsDao.AddGoods(obj)
	if err != nil {
		log.Fatalf("can't add goods from goodsDB service %v", err)
		return &rpc.AddResp{Success: false}, err
	}
	if ok == false {
		log.Fatalf("can't add goods, grpc ok")
		return &rpc.AddResp{Success: false}, nil
	}
	return &rpc.AddResp{Success: true}, nil
}

func (c *MasterServerImp) DelGoods(ctx context.Context, req *rpc.DelReq) (*rpc.DelResp, error) {
	ok, err := dao.GoodsDao.DelGoods(req.Id)
	if err != nil {
		log.Fatalf("can't delete goods from goodsDB service %v", err)
		return &rpc.DelResp{Success: false}, err
	}
	if ok == false {
		log.Fatalf("can't delete goods, grpc ok")
		return &rpc.DelResp{Success: false}, nil
	}
	return &rpc.DelResp{Success: true}, nil
}

func (c *MasterServerImp) UpdateGoods(ctx context.Context, req *rpc.UpdateReq) (*rpc.UpdateResp, error) {
	var obj *goodsModel.GoodsMod

	obj.Id = req.Id
	obj.Name = req.Name
	obj.Quantity = req.Quantity
	obj.Price = req.Price

	ok, err := dao.GoodsDao.UpdateGoods(req.Id, obj)
	if err != nil {
		log.Fatalf("can't update goods info %v", err)
		return &rpc.UpdateResp{Success: false}, err
	}

	if ok == false {
		log.Fatalf("can't update goods info, grpc ok")
		return &rpc.UpdateResp{Success: false}, nil
	}
	return &rpc.UpdateResp{Success: true}, nil
}

func (c *MasterServerImp) IfGoodsExist(ctx context.Context, req *rpc.IfExistReq) (*rpc.IfExistResp, error) {
	ok, err := dao.GoodsDao.GoodsExist(req.Id)
	if err != nil {
		log.Fatalf("can't check if goods exist or not %v", err)
		return &rpc.IfExistResp{Success: false}, err
	}
	if ok == false {
		log.Fatalf("goods does not exist")
		return &rpc.IfExistResp{Success: false}, nil
	}
	return &rpc.IfExistResp{Success: true}, nil
}
