package master

import (
	"log"
	rpc "project1_update/staffmana/Staffmanarpc"
	"project1_update/staffmana/microservice/staffDB"

	"golang.org/x/net/context"
)

type MasterServerImp struct{}

func (c *MasterServerImp) AddStaff(ctx context.Context, req *rpc.AddStaffReq) (*rpc.AddStaffResp, error) {
	r, err := staffDB.AddStaff(req.Name, req.Passwd, req.Position)
	if err != nil {
		log.Fatalf("can't add staff in master service of staffmana %v", err)
		return &rpc.AddStaffResp{Success: false}, err
	}
	return &rpc.AddStaffResp{Success: r}, err
}

func (c *MasterServerImp) DelStaff(ctx context.Context, req *rpc.DelStaffReq) (*rpc.DelStaffResp, error) {

	r, err := staffDB.DelStaff(req.Id)

	if err != nil {
		log.Fatalf("can't del staff from master server of staffmana %v", err)
		return &rpc.DelStaffResp{Success: false}, err
	}

	return &rpc.DelStaffResp{Success: r}, nil
}

func (c *MasterServerImp) UpdateStaff(ctx context.Context, req *rpc.UpdateStaffReq) (*rpc.UpdateStaffResp, error) {

	r, err := staffDB.UpdateStaff(req.Id, req.Name, req.Passwd, req.Position)

	if err != nil {
		log.Fatalf("can't update staff from master server of staffmana %v", err)
		return &rpc.UpdateStaffResp{Success: false}, err
	}

	return &rpc.UpdateStaffResp{Success: r}, nil
}

func (c *MasterServerImp) SelectAllStaff(ctx context.Context, req *rpc.SelectAllStaffReq) (*rpc.SelectAllStaffResp, error) {
	r, err := staffDB.SelectAllStaff()
	if err != nil {
		log.Fatalf("can't select all from staffmana %v", err)
		return &rpc.SelectAllStaffResp{Payload: nil}, err
	}
	return &rpc.SelectAllStaffResp{Payload: r}, err
}

func (c *MasterServerImp) SelectStaffById(ctx context.Context, req *rpc.SelectStaffByIdReq) (*rpc.SelectStaffByIdResp, error) {
	r, err := staffDB.SelectStaffById(req.Id)
	if err != nil {
		log.Fatalf("can't staff by id from staffmana %v", err)
		return &rpc.SelectStaffByIdResp{Payload: nil}, err
	}
	return &rpc.SelectStaffByIdResp{Payload: r}, err
}
