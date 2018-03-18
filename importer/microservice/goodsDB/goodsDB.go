package goodsDB

import (
	"log"
	"project1_update/importer/lib/MsgPack"
	Goods "project1_update/importer/microservice/services/GoodsDBrpc"
	GoodsModel "project1_update/others/models/goodsmod"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

func PricingAndUpdateGoodsQuantity(code []int, amount []int) (total int64, err error) {
	var goods *GoodsModel.GoodsMod
	conn, err := grpc.Dial("localhost:10001", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("importer conn to goodsDB err %v", err)
		return 0, err
	}
	defer conn.Close()

	c := Goods.NewGoodsClient(conn)
	for i := 0; i <= len(code); i++ {
		r, err := c.SelectByGoodsID(context.Background(), &Goods.SelectByIDReq{Id: int64(code[i])})
		if err != nil {
			log.Fatalf("can't select goods by id", err)
			return 0, err
		}
		err = MsgPack.Unmarshal(r.Payload, &goods)
		if err != nil {
			log.Fatalf("can't unmarshall in importer microservice %v", err)
			return 0, err
		}
		total = goods.Price * int64(amount[i])

		_, err1 := c.UpdateGoods(context.Background(), &Goods.UpdateReq{Id: int64(code[i]), Name: goods.Name, Quantity: goods.Quantity + int64(amount[i]), Price: goods.Price})
		if err1 != nil {
			log.Fatalf("can't update goods by id", err1)
			return 0, err1
		}
	}
	return total, nil
}

func Refund(code []int, amount []int) (ok bool, err error) {
	var goods *GoodsModel.GoodsMod
	conn, err := grpc.Dial("localhost:10001", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("exporter conn to goodsDB err %v", err)
		return false, err
	}
	defer conn.Close()

	c := Goods.NewGoodsClient(conn)
	for i := 0; i <= len(code); i++ {
		r, err := c.SelectByGoodsID(context.Background(), &Goods.SelectByIDReq{Id: int64(code[i])})
		if err != nil {
			log.Fatalf("can't select goods by id", err)
			return false, err
		}
		err = MsgPack.Unmarshal(r.Payload, &goods)
		if err != nil {
			log.Fatalf("can't unmarshall in exporter microservice %v", err)
			return false, err
		}
		_, err1 := c.UpdateGoods(context.Background(), &Goods.UpdateReq{Id: int64(code[i]), Name: goods.Name, Quantity: goods.Quantity - int64(amount[i]), Price: goods.Price})
		if err1 != nil {
			log.Fatalf("can't update goods by id", err1)
			return false, err1
		}
	}

	return true, nil

}
