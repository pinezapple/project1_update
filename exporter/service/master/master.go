package master

import (
	"log"
	rpc "project1_update/exporter/Exportrpc"
	"project1_update/exporter/dao"
	"project1_update/exporter/lib/MsgPack"
	"project1_update/exporter/microservice/goodsDB"
	DonModel "project1_update/others/models/reportmod"

	"golang.org/x/net/context"
)

type MasterServerImp struct{}

func (c *MasterServerImp) AddDE(ctx context.Context, req *rpc.AddDEReq) (*rpc.AddDEResp, error) {
	var obj *DonModel.DonExport
	var code, amount []int
	obj.Created_by = req.CreatedBy
	obj.Reason = req.Reason
	obj.Goods_list = req.GoodList
	obj.Goods_list_amount = req.GoodListAmount
	/*
		fo, err := os.Open(req.GoodList)
		if err != nil {
			log.Fatal("can't open the good_list in AddDE", err)
			return &rpc.AddDEResp{Success: false}, err
		}
		defer fo.Close()
		scanner := bufio.NewScanner(fo)
		scanner.Split(bufio.ScanLines)
		i := 0
		for scanner.Scan() {
			str := scanner.Text()
			size := strings.Split(str, " ")
			code1, err1 := strconv.Atoi(size[0])
			if err1 != nil {
				return &rpc.AddDEResp{Success: false}, err
			}
			amount1, err2 := strconv.Atoi(size[1])
			if err2 != nil {
				return &rpc.AddDEResp{Success: false}, err
			}
			code[i] = code1
			amount[i] = amount1
			i++
		}
	*/
	err := MsgPack.Unmarshal(obj.Goods_list, code)
	_ = MsgPack.Unmarshal(obj.Goods_list_amount, amount)
	obj.Total_money, err = goodsDB.PricingAndUpdateGoodsQuantity(code, amount)
	if err != nil {
		log.Fatalf("can't no calculate the money from the master service of exporter %v", err)
		return &rpc.AddDEResp{Success: false}, err
	}
	r, err := dao.ExporterDao.AddDE(obj)
	if err != nil {
		log.Fatalf("can't add to exporter db %v", err)
		return &rpc.AddDEResp{Success: false}, err
	}
	return &rpc.AddDEResp{Success: r}, nil
}

func (c *MasterServerImp) DelDE(ctx context.Context, req *rpc.DelDEReq) (*rpc.DelDEResp, error) {
	var code, amount []int
	r, goods_list, goods_list_amount, err := dao.ExporterDao.DelDE(req.Id)
	if err != nil {
		log.Fatalf("can't delete from exporter db %v", err)
		return &rpc.DelDEResp{Success: false}, err
	}
	_ = MsgPack.Unmarshal(goods_list, code)
	_ = MsgPack.Unmarshal(goods_list_amount, amount)
	_, _ = goodsDB.Refund(code, amount)
	return &rpc.DelDEResp{Success: r}, nil
}

func (c *MasterServerImp) UpdateDE(ctx context.Context, req *rpc.UpdateDEReq) (*rpc.UpdateDEResp, error) {
	var obj *DonModel.DonExport
	obj.Created_by = req.CreatedBy
	obj.Reason = req.Reason
	obj.Goods_list = req.GoodList
	obj.Total_money = req.TotalMoney
	r, err := dao.ExporterDao.UpdateDE(req.Id, obj)
	if err != nil {
		log.Fatalf("can't update in exporter db %v", err)
		return &rpc.UpdateDEResp{Success: false}, err
	}
	return &rpc.UpdateDEResp{Success: r}, nil
}

func (c *MasterServerImp) SelectAllDE(ctx context.Context, req *rpc.SelectAllDEReq) (*rpc.SelectAllDEResp, error) {
	var obj []*DonModel.DonExport
	obj, err := dao.ExporterDao.SelectAll()
	if err != nil {
		log.Fatalf("can't get all from exporter db %v", err)
		return &rpc.SelectAllDEResp{Payload: nil}, err
	}
	mar, err := MsgPack.Marshal(obj)
	if err != nil {
		log.Fatalf("can't marshall data in selectallde %v", err)
		return &rpc.SelectAllDEResp{Payload: nil}, err
	}
	return &rpc.SelectAllDEResp{Payload: mar}, nil
}

func (c *MasterServerImp) SelectByIDDE(ctx context.Context, req *rpc.SelectByIDDEReq) (*rpc.SelectByIDDEResp, error) {
	var obj *DonModel.DonExport
	obj, _, err := dao.ExporterDao.SelectByID(req.Id)
	if err != nil {
		log.Fatalf("can't get by id from exporter db %v", err)
		return &rpc.SelectByIDDEResp{Payload: nil}, err
	}
	mar, err := MsgPack.Marshal(obj)
	if err != nil {
		log.Fatalf("can't marshall data in selectbyidDE %v", err)
		return &rpc.SelectByIDDEResp{Payload: nil}, err
	}
	return &rpc.SelectByIDDEResp{Payload: mar}, nil
}

func (c *MasterServerImp) SelectByCreatedAtDE(ctx context.Context, req *rpc.SelectByCreatedAtDEReq) (*rpc.SelectByCreatedAtDEResp, error) {
	var obj []*DonModel.DonExport
	obj, _, err := dao.ExporterDao.SelectByCreatedAt(req.FromTime, req.ToTime)
	if err != nil {
		log.Fatalf("can't get by createdat from exporter db %v", err)
		return &rpc.SelectByCreatedAtDEResp{Payload: nil}, err
	}
	mar, err := MsgPack.Marshal(obj)
	if err != nil {
		log.Fatalf("can't marshall data in selectbycreatedatDE %v", err)
		return &rpc.SelectByCreatedAtDEResp{Payload: nil}, err
	}
	return &rpc.SelectByCreatedAtDEResp{Payload: mar}, nil
}
