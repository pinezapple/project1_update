package master

import (
	"log"
	rpc "project1_update/authen/Authenrpc"
	"project1_update/authen/dao"
	"project1_update/authen/lib/MsgPack"
	"project1_update/authen/microservice/staffDB"
	DonModel "project1_update/others/models/reportmod"

	"golang.org/x/net/context"
)

type MasterServerImp struct{}

func (c *MasterServerImp) AuthenAndClasify(ctx context.Context, req *rpc.AuthenAndClasifyReq) (*rpc.AuthenAndClasifyResp, error) {
	r, err := staffDB.AuthenAndClassify(req.Staffid, req.Passwd, req.Position)
	if err != nil {
		log.Fatalf("can't authen in master service of authen %v", err)
		return &rpc.AuthenAndClasifyResp{Ok: false}, err
	}
	return &rpc.AuthenAndClasifyResp{Ok: r}, err
}

func (c *MasterServerImp) AddHour(ctx context.Context, req *rpc.AddHourReq) (*rpc.AddHourResp, error) {

	var obj *DonModel.HoursByDay
	obj.StaffId = req.Staffid
	obj.In = req.In
	obj.Out = req.Out

	r, err := dao.HoursByDay.AddHour(obj)

	if err != nil {
		log.Fatalf("can't add hour from master server of authen %v", err)
		return &rpc.AddHourResp{Success: false}, err
	}

	return &rpc.AddHourResp{Success: r}, nil
}

func (c *MasterServerImp) SelectAllHour(ctx context.Context, req *rpc.SelectAllHourReq) (*rpc.SelectAllHourResp, error) {
	r, err := dao.HoursByDay.SelectAll()
	if err != nil {
		log.Fatalf("can't select all from authen_db %v", err)
		return &rpc.SelectAllHourResp{Payload: nil}, err
	}
	mar, err := MsgPack.Marshal(r)
	if err != nil {
		log.Fatalf("can't not marshall all hours %v", err)
		return &rpc.SelectAllHourResp{Payload: nil}, err
	}
	return &rpc.SelectAllHourResp{Payload: mar}, err
}

func (c *MasterServerImp) SelectHourByStaffID(ctx context.Context, req *rpc.SelectHourByStaffIDReq) (*rpc.SelectHourByStaffIDResp, error) {
	r, _, err := dao.HoursByDay.SelectByStaffID(req.Staffid)
	if err != nil {
		log.Fatalf("can't select all from authen_db %v", err)
		return &rpc.SelectHourByStaffIDResp{Payload: nil}, err
	}
	mar, err := MsgPack.Marshal(r)
	if err != nil {
		log.Fatalf("can't not marshall all hours %v", err)
		return &rpc.SelectHourByStaffIDResp{Payload: nil}, err
	}
	return &rpc.SelectHourByStaffIDResp{Payload: mar}, err
}
