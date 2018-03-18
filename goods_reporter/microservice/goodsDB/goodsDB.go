package goodsDB

import (
	"log"
	Goods "project1_update/goods_reporter/microservice/services/GoodsDBrpc"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

func IfGoodExist(id int64) (ok bool, err error) {

	conn, err := grpc.Dial("localhost:10001", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("goods_reporter conn to goodsDB err %v", err)
		return false, err
	}
	defer conn.Close()

	c := Goods.NewGoodsClient(conn)

	r, err := c.IfGoodsExist(context.Background(), &Goods.IfExistReq{Id: id})

	if err != nil {
		log.Fatalf("can't check for goods in good_reporter %v", err)
		return false, err
	}

	return r.Success, nil
}

func SelectAllGoods() (mar []byte, err error) {
	conn, err := grpc.Dial("localhost:10001", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("goods_reporter conn to goodsDB err %v", err)
		return nil, err
	}
	defer conn.Close()

	c := Goods.NewGoodsClient(conn)

	r, err := c.SelectAllGoods(context.Background(), &Goods.SelectAllReq{})

	if err != nil {
		log.Fatalf("can't search for all goods in good_reporter %v", err)
		return nil, err
	}

	return r.Payload, nil
}

func SelectGoodsById(id int64) (mar []byte, err error) {
	conn, err := grpc.Dial("localhost:10001", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("goods_reporter conn to goodsDB err %v", err)
		return nil, err
	}
	defer conn.Close()

	c := Goods.NewGoodsClient(conn)

	r, err := c.SelectByGoodsID(context.Background(), &Goods.SelectByIDReq{Id: id})

	if err != nil {
		log.Fatalf("can't search good by id in good_reporter %v", err)
		return nil, err
	}

	return r.Payload, nil
}

func SelectGoodsByName(name string) (mar []byte, err error) {
	conn, err := grpc.Dial("localhost:10001", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("goods_reporter conn to goodsDB err %v", err)
		return nil, err
	}
	defer conn.Close()

	c := Goods.NewGoodsClient(conn)

	r, err := c.SelectByGoodsName(context.Background(), &Goods.SelectByNameReq{Name: name})

	if err != nil {
		log.Fatalf("can't search for goods by name in good_reporter %v", err)
		return nil, err
	}

	return r.Payload, nil
}

func SelectGoodsByPrice(price int64) (mar []byte, err error) {
	conn, err := grpc.Dial("localhost:10001", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("goods_reporter conn to goodsDB err %v", err)
		return nil, err
	}
	defer conn.Close()

	c := Goods.NewGoodsClient(conn)

	r, err := c.SelectByGoodsPrice(context.Background(), &Goods.SelectByGoodsPriceReq{Price: price})

	if err != nil {
		log.Fatalf("can't search for goods by price in good_reporter %v", err)
		return nil, err
	}

	return r.Payload, nil
}
