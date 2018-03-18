package master

import (
	"log"
	rpc "project1_update/staffDB/StaffDBrpc"
	"project1_update/staffDB/dao"
	"project1_update/staffDB/lib/MsgPack"

	staffModel "project1_update/others/models/staffmod"

	"golang.org/x/net/context"
)

type MasterServerImp struct{}

func (c *MasterServerImp) SelectAllStaff(ctx context.Context, req *rpc.SelectAllReq) (*rpc.SelectAllResp, error) {
	r, err := dao.StaffDao.SelectAll()
	if err != nil {
		log.Fatalf("can't select all staff from staffDB service %v", err)
		return &rpc.SelectAllResp{Payload: nil}, err
	}
	mar, _ := MsgPack.Marshal(r)
	return &rpc.SelectAllResp{Payload: mar}, nil
}

func (c *MasterServerImp) SelectByStaffID(ctx context.Context, req *rpc.SelectByIDReq) (*rpc.SelectByIDResp, error) {
	r, ok, err := dao.StaffDao.SelectByID(req.Id)
	if err != nil {
		log.Fatalf("can't select staff by id from staffDB service %v", err)
		return &rpc.SelectByIDResp{Payload: nil}, err
	}
	if ok == false {
		log.Fatalf("staff does not exist")
		return &rpc.SelectByIDResp{Payload: nil}, nil
	}
	mar, _ := MsgPack.Marshal(r)
	return &rpc.SelectByIDResp{Payload: mar}, nil
}

func (c *MasterServerImp) AddStaff(ctx context.Context, req *rpc.AddReq) (*rpc.AddResp, error) {
	var obj *staffModel.StaffMod

	obj.Name = req.Name
	obj.Passwd = req.Passwd
	obj.Position = req.Position

	ok, err := dao.StaffDao.AddStaff(obj)
	if err != nil {
		log.Fatalf("can't add staff from staffDB service %v", err)
		return &rpc.AddResp{Success: false}, err
	}
	if ok == false {
		log.Fatalf("can't add staff, grpc ok")
		return &rpc.AddResp{Success: false}, nil
	}
	return &rpc.AddResp{Success: true}, nil
}

func (c *MasterServerImp) DelStaff(ctx context.Context, req *rpc.DelReq) (*rpc.DelResp, error) {
	ok, err := dao.StaffDao.DelStaff(req.Id)
	if err != nil {
		log.Fatalf("can't delete staff from staffDB service %v", err)
		return &rpc.DelResp{Success: false}, err
	}
	if ok == false {
		log.Fatalf("can't delete staff, grpc ok")
		return &rpc.DelResp{Success: false}, nil
	}
	return &rpc.DelResp{Success: true}, nil
}

func (c *MasterServerImp) UpdateStaff(ctx context.Context, req *rpc.UpdateReq) (*rpc.UpdateResp, error) {
	var obj *staffModel.StaffMod

	obj.Id = req.Id
	obj.Name = req.Name
	obj.Passwd = req.Passwd
	obj.Position = req.Position

	ok, err := dao.StaffDao.UpdateStaff(req.Id, obj)
	if err != nil {
		log.Fatalf("can't update staff info %v", err)
		return &rpc.UpdateResp{Success: false}, err
	}

	if ok == false {
		log.Fatalf("can't update staff info, grpc ok")
		return &rpc.UpdateResp{Success: false}, nil
	}
	return &rpc.UpdateResp{Success: true}, nil
}

func (c *MasterServerImp) IfStaffExist(ctx context.Context, req *rpc.IfExistReq) (*rpc.IfExistResp, error) {
	ok, err := dao.StaffDao.StaffExist(req.Id)
	if err != nil {
		log.Fatalf("can't check if staff exist or not %v", err)
		return &rpc.IfExistResp{Ok: false}, err
	}
	if ok == false {
		log.Fatalf("staff does not exist")
		return &rpc.IfExistResp{Ok: false}, nil
	}
	return &rpc.IfExistResp{Ok: true}, nil
}
