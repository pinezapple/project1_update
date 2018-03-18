package master

import (
	"log"
	rpc "project1_update/importer/Importerrpc"
	"project1_update/importer/dao"
	"project1_update/importer/lib/MsgPack"
	"project1_update/importer/microservice/goodsDB"
	DonModel "project1_update/others/models/reportmod"

	"golang.org/x/net/context"
)

type MasterServerImp struct{}

func (c *MasterServerImp) AddDI(ctx context.Context, req *rpc.AddDIReq) (*rpc.AddDIResp, error) {
	var obj *DonModel.DonImport

	var code, amount []int
	obj.Created_by = req.CreatedBy
	obj.Goods_list = req.GoodList
	obj.Goods_list_amount = req.GoodListAmount

	err := MsgPack.Unmarshal(obj.Goods_list, code)

	_ = MsgPack.Unmarshal(obj.Goods_list_amount, amount)
	/*
		fo, err := os.Open(req.GoodList)
		if err != nil {
			log.Fatal("can't open the good_list in AddDI", err)
			return &rpc.AddDIResp{Success: false}, err
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
					return &rpc.AddDIResp{Success: false}, err
				}
				amount1, err2 := strconv.Atoi(size[1])
				if err2 != nil {
					return &rpc.AddDIResp{Success: false}, err
				}
				code[i] = code1
				amount[i] = amount1
				i++
			}
	*/

	obj.Total_money, err = goodsDB.PricingAndUpdateGoodsQuantity(code, amount)
	if err != nil {
		log.Fatalf("can't no calculate the money from the master service of importer %v", err)
		return &rpc.AddDIResp{Success: false}, err
	}
	r, err := dao.ImporterDao.AddDI(obj)
	if err != nil {
		log.Fatalf("can't add to importer db %v", err)
		return &rpc.AddDIResp{Success: false}, err
	}
	return &rpc.AddDIResp{Success: r}, nil
}

func (c *MasterServerImp) DelDI(ctx context.Context, req *rpc.DelDIReq) (*rpc.DelDIResp, error) {
	var code, amount []int
	r, goods_list, good_list_amount, err := dao.ImporterDao.DelDI(req.Id)
	_ = MsgPack.Unmarshal(goods_list, code)
	_ = MsgPack.Unmarshal(good_list_amount, amount)
	_, _ = goodsDB.Refund(code, amount)
	if err != nil {
		log.Fatalf("can't delete from importer db %v", err)
		return &rpc.DelDIResp{Success: false}, err
	}

	return &rpc.DelDIResp{Success: r}, nil
}

func (c *MasterServerImp) UpdateDI(ctx context.Context, req *rpc.UpdateDIReq) (*rpc.UpdateDIResp, error) {
	var obj *DonModel.DonImport
	var code, amount []int
	var err error
	obj.Created_by = req.CreatedBy
	obj.Goods_list = req.GoodList
	obj.Goods_list_amount = req.GoodListAmount

	obj.Total_money, err = goodsDB.PricingAndUpdateGoodsQuantity(code, amount)
	if err != nil {
		log.Fatalf("can't no calculate the money from the master service of importer %v", err)
		return &rpc.UpdateDIResp{Success: false}, err
	}

	r, err := dao.ImporterDao.UpdateDI(req.Id, obj)
	if err != nil {
		log.Fatalf("can't update in importer db %v", err)
		return &rpc.UpdateDIResp{Success: false}, err
	}
	return &rpc.UpdateDIResp{Success: r}, nil
}

func (c *MasterServerImp) SelectAllDI(ctx context.Context, req *rpc.SelectAllDIReq) (*rpc.SelectAllDIResp, error) {
	var obj []*DonModel.DonImport
	obj, err := dao.ImporterDao.SelectAll()
	if err != nil {
		log.Fatalf("can't get all from importer db %v", err)
		return &rpc.SelectAllDIResp{Payload: nil}, err
	}
	mar, err := MsgPack.Marshal(obj)
	if err != nil {
		log.Fatalf("can't marshall data in selectalldi %v", err)
		return &rpc.SelectAllDIResp{Payload: nil}, err
	}
	return &rpc.SelectAllDIResp{Payload: mar}, nil
}

func (c *MasterServerImp) SelectByIDDI(ctx context.Context, req *rpc.SelectByIDDIReq) (*rpc.SelectByIDDIResp, error) {
	var obj *DonModel.DonImport
	obj, _, err := dao.ImporterDao.SelectByID(req.Id)
	if err != nil {
		log.Fatalf("can't get by id from importer db %v", err)
		return &rpc.SelectByIDDIResp{Payload: nil}, err
	}
	mar, err := MsgPack.Marshal(obj)
	if err != nil {
		log.Fatalf("can't marshall data in selectbyidDI %v", err)
		return &rpc.SelectByIDDIResp{Payload: nil}, err
	}
	return &rpc.SelectByIDDIResp{Payload: mar}, nil
}

func (c *MasterServerImp) SelectByCreatedAtDI(ctx context.Context, req *rpc.SelectByCreatedAtDIReq) (*rpc.SelectByCreatedAtDIResp, error) {
	var obj []*DonModel.DonImport
	obj, _, err := dao.ImporterDao.SelectByCreatedAt(req.FromTime, req.ToTime)
	if err != nil {
		log.Fatalf("can't get by createdat from importer db %v", err)
		return &rpc.SelectByCreatedAtDIResp{Payload: nil}, err
	}
	mar, err := MsgPack.Marshal(obj)
	if err != nil {
		log.Fatalf("can't marshall data in selectbycreatedatDI %v", err)
		return &rpc.SelectByCreatedAtDIResp{Payload: nil}, err
	}
	return &rpc.SelectByCreatedAtDIResp{Payload: mar}, nil
}
