package goods_reporter

import (
	"log"
	Goods "project1_update/boss_user/microservice/goods_reporter/Goodreportrpc"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

func IfGoodExist(id int64) (ok bool, err error) {

	conn, err := grpc.Dial("localhost:10009", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user conn to good reporter err %v", err)
		return false, err
	}
	defer conn.Close()

	c := Goods.NewGoodreportClient(conn)

	r, err := c.IfGoodExist(context.Background(), &Goods.IfGoodExistReq{Id: id})

	if err != nil {
		log.Fatalf("can't check for goods in boss user %v", err)
		return false, err
	}

	return r.Ok, nil
}

func SelectAllGoods() (mar []byte, err error) {
	conn, err := grpc.Dial("localhost:10009", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user conn to good reporter err %v", err)
		return nil, err
	}
	defer conn.Close()

	c := Goods.NewGoodreportClient(conn)

	r, err := c.SelectAllGood(context.Background(), &Goods.SelectAllGoodReq{})

	if err != nil {
		log.Fatalf("can't search for all goods in boss user %v", err)
		return nil, err
	}

	return r.Payload, nil
}

func SelectGoodsById(id int64) (mar []byte, err error) {
	conn, err := grpc.Dial("localhost:10009", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user conn to good reporter err %v", err)
		return nil, err
	}
	defer conn.Close()

	c := Goods.NewGoodreportClient(conn)

	r, err := c.SelectGoodById(context.Background(), &Goods.SelectGoodByIdReq{Id: id})

	if err != nil {
		log.Fatalf("can't search good by id in boss user %v", err)
		return nil, err
	}

	return r.Payload, nil
}

func SelectGoodsByName(name string) (mar []byte, err error) {
	conn, err := grpc.Dial("localhost:10009", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user conn to good reporter err %v", err)
		return nil, err
	}
	defer conn.Close()

	c := Goods.NewGoodreportClient(conn)

	r, err := c.SelectGoodByName(context.Background(), &Goods.SelectGoodByNameReq{Name: name})

	if err != nil {
		log.Fatalf("can't search for goods by name in boss user %v", err)
		return nil, err
	}

	return r.Payload, nil
}

func SelectGoodsByPrice(price int64) (mar []byte, err error) {
	conn, err := grpc.Dial("localhost:10009", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user conn to good reporter err %v", err)
		return nil, err
	}
	defer conn.Close()

	c := Goods.NewGoodreportClient(conn)

	r, err := c.SelectGoodByPrice(context.Background(), &Goods.SelectGoodByPriceReq{Price: price})

	if err != nil {
		log.Fatalf("can't search for goods by price in boss user %v", err)
		return nil, err
	}

	return r.Payload, nil
}
