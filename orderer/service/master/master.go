package master

import (
	"log"
	rpc "project1_update/orderer/Ordererrpc"
	"project1_update/orderer/dao"
	"project1_update/orderer/lib/MsgPack"
	"project1_update/orderer/microservice/customerDB"
	"project1_update/orderer/microservice/goodsDB"
	DonModel "project1_update/others/models/reportmod"
	"strconv"

	"golang.org/x/net/context"
)

type MasterServerImp struct{}

func (c *MasterServerImp) AddDS(ctx context.Context, req *rpc.AddDSReq) (*rpc.AddDSResp, error) {
	var obj *DonModel.DonSell
	var code, amount []int

	obj.Seller_name = req.SellerName
	obj.Cus_name = req.CusName
	obj.Goods_list = req.GoodList
	obj.Goods_list_amount = req.GoodListAmount

	err := MsgPack.Unmarshal(obj.Goods_list, code)
	_ = MsgPack.Unmarshal(obj.Goods_list_amount, amount)
	/*
		fo, err := os.Open(req.GoodList)
		if err != nil {
			log.Fatal("can't open the good_list in AddDS", err)
			return &rpc.AddDSResp{Success: false}, err
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
				return &rpc.AddDSResp{Success: false}, err
			}
			amount1, err2 := strconv.Atoi(size[1])
			if err2 != nil {
				return &rpc.AddDSResp{Success: false}, err
			}
			code[i] = code1
			amount[i] = amount1
			i++
		}
	*/

	obj.Total_money, err = goodsDB.PricingAndUpdateGoodsQuantity(code, amount)
	n, _ := strconv.ParseInt(req.CusName, 10, 64)

	_, err3 := customerDB.AddToCusBalance(n, obj.Total_money)

	if err3 != nil {
		log.Fatalf("can't add to cus balance %v", err)
		return &rpc.AddDSResp{Success: false}, err
	}
	if err != nil {
		log.Fatalf("can't no calculate the money from the master service of orderer %v", err)
		return &rpc.AddDSResp{Success: false}, err
	}
	r, err := dao.OrdererDao.AddDS(obj)
	if err != nil {
		log.Fatalf("can't add to orderer db %v", err)
		return &rpc.AddDSResp{Success: false}, err
	}
	return &rpc.AddDSResp{Success: r}, nil
}

func (c *MasterServerImp) DelDS(ctx context.Context, req *rpc.DelDSReq) (*rpc.DelDSResp, error) {
	var code, amount []int
	r, total_money, cusname, goods_list, goods_list_amount, err := dao.OrdererDao.DelDS(req.Id)
	if err != nil {
		log.Fatalf("can't delete from orderer db %v", err)
		return &rpc.DelDSResp{Success: false}, err
	}
	cus_name, _ := strconv.ParseInt(cusname, 10, 64)
	_ = MsgPack.Unmarshal(goods_list, code)
	_ = MsgPack.Unmarshal(goods_list_amount, amount)
	_, _ = goodsDB.Refund(code, amount)
	_, _ = customerDB.ResetBalance(cus_name, total_money)
	return &rpc.DelDSResp{Success: r}, nil
}

func (c *MasterServerImp) UpdateDS(ctx context.Context, req *rpc.UpdateDSReq) (*rpc.UpdateDSResp, error) {
	var obj *DonModel.DonSell

	obj.Seller_name = req.SellerName
	obj.Cus_name = req.CusName
	obj.Goods_list = req.GoodList
	obj.Total_money = req.TotalMoney

	r, err := dao.OrdererDao.UpdateDS(req.Id, obj)
	if err != nil {
		log.Fatalf("can't update in orderer db %v", err)
		return &rpc.UpdateDSResp{Success: false}, err
	}
	return &rpc.UpdateDSResp{Success: r}, nil
}

func (c *MasterServerImp) SelectAllDS(ctx context.Context, req *rpc.SelectAllDSReq) (*rpc.SelectAllDSResp, error) {
	var obj []*DonModel.DonSell
	obj, err := dao.OrdererDao.SelectAll()
	if err != nil {
		log.Fatalf("can't get all from orderer db %v", err)
		return &rpc.SelectAllDSResp{Payload: nil}, err
	}
	mar, err := MsgPack.Marshal(obj)
	if err != nil {
		log.Fatalf("can't marshall data in selectallds %v", err)
		return &rpc.SelectAllDSResp{Payload: nil}, err
	}
	return &rpc.SelectAllDSResp{Payload: mar}, nil
}

func (c *MasterServerImp) SelectByIDDS(ctx context.Context, req *rpc.SelectByIDDSReq) (*rpc.SelectByIDDSResp, error) {
	var obj *DonModel.DonSell
	obj, _, err := dao.OrdererDao.SelectByID(req.Id)
	if err != nil {
		log.Fatalf("can't get by id from orderer db %v", err)
		return &rpc.SelectByIDDSResp{Payload: nil}, err
	}
	mar, err := MsgPack.Marshal(obj)
	if err != nil {
		log.Fatalf("can't marshall data in selectbyidDS %v", err)
		return &rpc.SelectByIDDSResp{Payload: nil}, err
	}
	return &rpc.SelectByIDDSResp{Payload: mar}, nil
}

func (c *MasterServerImp) SelectByCreatedAtDS(ctx context.Context, req *rpc.SelectByCreatedAtDSReq) (*rpc.SelectByCreatedAtDSResp, error) {
	var obj []*DonModel.DonSell
	obj, _, err := dao.OrdererDao.SelectByCreatedAt(req.FromTime, req.ToTime)
	if err != nil {
		log.Fatalf("can't get by createdat from orderer db %v", err)
		return &rpc.SelectByCreatedAtDSResp{Payload: nil}, err
	}
	mar, err := MsgPack.Marshal(obj)
	if err != nil {
		log.Fatalf("can't marshall data in selectbycreatedatDS %v", err)
		return &rpc.SelectByCreatedAtDSResp{Payload: nil}, err
	}
	return &rpc.SelectByCreatedAtDSResp{Payload: mar}, nil
}
