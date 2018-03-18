package master

import (
	"log"
	rpc "project1_update/spender/Spenderrpc"
	"project1_update/spender/dao"
	"project1_update/spender/lib/MsgPack"

	DonModel "project1_update/others/models/reportmod"

	"golang.org/x/net/context"
)

type MasterServerImp struct{}

func (c *MasterServerImp) AddDSp(ctx context.Context, req *rpc.AddDSpReq) (*rpc.AddDSpResp, error) {
	var obj *DonModel.DonSpend

	obj.Created_by = req.CreatedBy
	obj.Reason = req.Reason
	obj.Total_money = req.TotalMoney

	r, err := dao.SpenderDao.AddDSp(obj)
	if err != nil {
		log.Fatalf("can't add to Spender db %v", err)
		return &rpc.AddDSpResp{Success: false}, err
	}
	return &rpc.AddDSpResp{Success: r}, nil
}

func (c *MasterServerImp) DelDSp(ctx context.Context, req *rpc.DelDSpReq) (*rpc.DelDSpResp, error) {
	r, err := dao.SpenderDao.DelDSp(req.Id)
	if err != nil {
		log.Fatalf("can't delete from spender db %v", err)
		return &rpc.DelDSpResp{Success: false}, err
	}

	return &rpc.DelDSpResp{Success: r}, nil
}

func (c *MasterServerImp) UpdateDSp(ctx context.Context, req *rpc.UpdateDSpReq) (*rpc.UpdateDSpResp, error) {
	var obj *DonModel.DonSpend

	obj.Created_by = req.CreatedBy
	obj.Reason = req.Reason
	obj.Total_money = req.TotalMoney

	r, err := dao.SpenderDao.UpdateDSp(req.Id, obj)
	if err != nil {
		log.Fatalf("can't update in spender db %v", err)
		return &rpc.UpdateDSpResp{Success: false}, err
	}
	return &rpc.UpdateDSpResp{Success: r}, nil
}

func (c *MasterServerImp) SelectAllDSp(ctx context.Context, req *rpc.SelectAllDSpReq) (*rpc.SelectAllDSpResp, error) {
	var obj []*DonModel.DonSpend
	obj, err := dao.SpenderDao.SelectAll()
	if err != nil {
		log.Fatalf("can't get all from spender db %v", err)
		return &rpc.SelectAllDSpResp{Payload: nil}, err
	}
	mar, err := MsgPack.Marshal(obj)
	if err != nil {
		log.Fatalf("can't marshall data in selectalldsp %v", err)
		return &rpc.SelectAllDSpResp{Payload: nil}, err
	}
	return &rpc.SelectAllDSpResp{Payload: mar}, nil
}

func (c *MasterServerImp) SelectByIDDSp(ctx context.Context, req *rpc.SelectByIDDSpReq) (*rpc.SelectByIDDSpResp, error) {
	var obj *DonModel.DonSpend
	obj, _, err := dao.SpenderDao.SelectByID(req.Id)
	if err != nil {
		log.Fatalf("can't get by id from spender db %v", err)
		return &rpc.SelectByIDDSpResp{Payload: nil}, err
	}
	mar, err := MsgPack.Marshal(obj)
	if err != nil {
		log.Fatalf("can't marshall data in selectbyidDSp %v", err)
		return &rpc.SelectByIDDSpResp{Payload: nil}, err
	}
	return &rpc.SelectByIDDSpResp{Payload: mar}, nil
}

func (c *MasterServerImp) SelectByCreatedAtDSp(ctx context.Context, req *rpc.SelectByCreatedAtDSpReq) (*rpc.SelectByCreatedAtDSpResp, error) {
	var obj []*DonModel.DonSpend
	obj, _, err := dao.SpenderDao.SelectByCreatedAt(req.FromTime, req.ToTime)
	if err != nil {
		log.Fatalf("can't get by createdat from spender db %v", err)
		return &rpc.SelectByCreatedAtDSpResp{Payload: nil}, err
	}
	mar, err := MsgPack.Marshal(obj)
	if err != nil {
		log.Fatalf("can't marshall data in selectbycreatedatDSp %v", err)
		return &rpc.SelectByCreatedAtDSpResp{Payload: nil}, err
	}
	return &rpc.SelectByCreatedAtDSpResp{Payload: mar}, nil
}
