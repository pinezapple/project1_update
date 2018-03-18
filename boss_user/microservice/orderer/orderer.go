package orderer

import (
	"log"
	Orderer "project1_update/boss_user/microservice/orderer/Ordererrpc"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

func AddDS(seller_name string, cus_name string, goods_list []byte, goods_list_amount []byte) (ok bool, err error) {
	conn, err := grpc.Dial("localhost:10006", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to orderer grpc connection fail %v", err)
		return false, err
	}
	defer conn.Close()

	c := Orderer.NewOrdererClient(conn)
	r, err := c.AddDS(context.Background(), &Orderer.AddDSReq{SellerName: seller_name, CusName: cus_name, GoodList: goods_list, GoodListAmount: goods_list_amount, TotalMoney: 0})
	if err != nil {
		log.Fatalf("Can't Add Don Order from boss user", err)
		return false, err
	}
	return r.Success, nil
}

func DelDS(id int64) (ok bool, err error) {
	conn, err := grpc.Dial("localhost:10006", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to orderer grpc connection fail %v", err)
		return false, err
	}
	defer conn.Close()

	c := Orderer.NewOrdererClient(conn)
	r, err := c.DelDS(context.Background(), &Orderer.DelDSReq{Id: id})
	if err != nil {
		log.Fatalf("can't delete Don Order from boss user", err)
		return false, err
	}
	return r.Success, nil
}

func SelectAllDS() ([]byte, error) {
	conn, err := grpc.Dial("localhost:10006", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to orderer grpc connection fail %v", err)
		return nil, err
	}
	defer conn.Close()

	c := Orderer.NewOrdererClient(conn)
	r, err := c.SelectAllDS(context.Background(), &Orderer.SelectAllDSReq{})
	if err != nil {
		log.Fatalf("can't select all Don Order from boss user", err)
		return nil, err
	}
	return r.Payload, nil
}

func SelectByIDDS(id int64) ([]byte, error) {
	conn, err := grpc.Dial("localhost:10006", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to orderer grpc connection fail %v", err)
		return nil, err
	}
	defer conn.Close()

	c := Orderer.NewOrdererClient(conn)
	r, err := c.SelectByIDDS(context.Background(), &Orderer.SelectByIDDSReq{Id: id})
	if err != nil {
		log.Fatalf("can't select Don orderer by id from boss user", err)
		return nil, err
	}
	return r.Payload, nil
}

func SelectByCreatedAtDS(from_time int64, to_time int64) ([]byte, error) {
	conn, err := grpc.Dial("localhost:10006", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to orderer grpc connection fail %v", err)
		return nil, err
	}
	defer conn.Close()

	c := Orderer.NewOrdererClient(conn)
	r, err := c.SelectByCreatedAtDS(context.Background(), &Orderer.SelectByCreatedAtDSReq{FromTime: from_time, ToTime: to_time})
	if err != nil {
		log.Fatalf("can't select Don order by created at from boss user", err)
		return nil, err
	}
	return r.Payload, nil
}
