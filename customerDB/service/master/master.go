package master

import (
	"log"
	rpc "project1_update/customerDB/CustomerDBrpc"
	"project1_update/customerDB/dao"
	"project1_update/customerDB/lib/MsgPack"

	cusModel "project1_update/others/models/customermod"

	"golang.org/x/net/context"
)

type MasterServerImp struct{}

func (c *MasterServerImp) SelectAllCus(ctx context.Context, req *rpc.SelectAllReq) (*rpc.SelectAllResp, error) {
	r, err := dao.CusDao.SelectAll()
	if err != nil {
		log.Fatalf("can't select all customers from cusDB service %v", err)
		return &rpc.SelectAllResp{Payload: nil}, err
	}
	mar, _ := MsgPack.Marshal(r)
	return &rpc.SelectAllResp{Payload: mar}, nil
}

func (c *MasterServerImp) SelectByCusID(ctx context.Context, req *rpc.SelectByIDReq) (*rpc.SelectByIDResp, error) {
	r, ok, err := dao.CusDao.SelectByID(req.Id)
	if err != nil {
		log.Fatalf("can't select customers by id from cusDB service %v", err)
		return &rpc.SelectByIDResp{Payload: nil}, err
	}
	if ok == false {
		log.Fatalf("customer does not exist")
		return &rpc.SelectByIDResp{Payload: nil}, nil
	}
	mar, _ := MsgPack.Marshal(r)
	return &rpc.SelectByIDResp{Payload: mar}, nil
}

func (c *MasterServerImp) SelectByCusName(ctx context.Context, req *rpc.SelectByNameReq) (*rpc.SelectByNameResp, error) {
	r, ok, err := dao.CusDao.SelectByName(req.Name)
	if err != nil {
		log.Fatalf("can't select customers by name from cusDB service %v", err)
		return &rpc.SelectByNameResp{Payload: nil}, err
	}
	if ok == false {
		log.Fatalf("customer does not exist")
		return &rpc.SelectByNameResp{Payload: nil}, nil
	}
	mar, _ := MsgPack.Marshal(r)
	return &rpc.SelectByNameResp{Payload: mar}, nil
}

func (c *MasterServerImp) SelectByCusPhoneNum(ctx context.Context, req *rpc.SelectByPhoneNumReq) (*rpc.SelectByPhoneNumResp, error) {
	r, ok, err := dao.CusDao.SelectByPhoneNum(req.Phonenum)
	if err != nil {
		log.Fatalf("can't select customers by name from cusDB service %v", err)
		return &rpc.SelectByPhoneNumResp{Payload: nil}, err
	}
	if ok == false {
		log.Fatalf("customer does not exist")
		return &rpc.SelectByPhoneNumResp{Payload: nil}, nil
	}
	mar, _ := MsgPack.Marshal(r)
	return &rpc.SelectByPhoneNumResp{Payload: mar}, nil
}

func (c *MasterServerImp) AddCus(ctx context.Context, req *rpc.AddReq) (*rpc.AddResp, error) {
	var obj *cusModel.CustomerMod

	obj.Id = req.Id
	obj.Balance = req.Balance
	obj.Name = req.Name
	obj.Phonenum = req.Phonenum
	obj.Level = req.Level

	ok, err := dao.CusDao.AddCus(obj)
	if err != nil {
		log.Fatalf("can't add customer from cusDB service %v", err)
		return &rpc.AddResp{Success: false}, err
	}
	if ok == false {
		log.Fatalf("can't add customer, grpc ok")
		return &rpc.AddResp{Success: false}, nil
	}
	return &rpc.AddResp{Success: true}, nil
}

func (c *MasterServerImp) DelCus(ctx context.Context, req *rpc.DelReq) (*rpc.DelResp, error) {
	ok, err := dao.CusDao.DelCus(req.Id)
	if err != nil {
		log.Fatalf("can't del customer from cusDB service %v", err)
		return &rpc.DelResp{Success: false}, err
	}
	if ok == false {
		log.Fatalf("can't delete customer, grpc ok")
		return &rpc.DelResp{Success: false}, nil
	}
	return &rpc.DelResp{Success: true}, nil
}

func (c *MasterServerImp) UpdateCus(ctx context.Context, req *rpc.UpdateReq) (*rpc.UpdateResp, error) {
	var obj *cusModel.CustomerMod

	obj.Id = req.Id
	obj.Balance = req.Balance
	obj.Name = req.Name
	obj.Phonenum = req.Phonenum
	obj.Level = req.Level

	ok, err := dao.CusDao.UpdateCus(req.Id, obj)
	if err != nil {
		log.Fatalf("can't update cus info %v", err)
		return &rpc.UpdateResp{Success: false}, err
	}

	if ok == false {
		log.Fatalf("can't update cus info, grpc ok")
		return &rpc.UpdateResp{Success: false}, nil
	}
	return &rpc.UpdateResp{Success: true}, nil
}

func (c *MasterServerImp) AddToCusBalance(ctx context.Context, req *rpc.AddToBalanceReq) (*rpc.AddToBalanceResp, error) {
	ok, err := dao.CusDao.AddToBalance(req.Id, req.Amount)
	if err != nil {
		log.Fatalf("can't add to cus balance %v", err)
		return &rpc.AddToBalanceResp{Success: false}, err
	}
	if ok == false {
		log.Fatalf("can't add to cus balance, grpc ok")
		return &rpc.AddToBalanceResp{Success: false}, nil
	}
	return &rpc.AddToBalanceResp{Success: true}, nil
}

func (c *MasterServerImp) IfCusExist(ctx context.Context, req *rpc.CusExistReq) (*rpc.CusExistResp, error) {
	ok, err := dao.CusDao.CusExist(req.Id)
	if err != nil {
		log.Fatalf("can't check if cus exist or not %v", err)
		return &rpc.CusExistResp{Ok: false}, err
	}
	if ok == false {
		log.Fatalf("cus does not exist")
		return &rpc.CusExistResp{Ok: false}, nil
	}
	return &rpc.CusExistResp{Ok: true}, nil
}
