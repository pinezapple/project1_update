package exporter

import (
	"log"
	Exporter "project1_update/boss_user/microservice/exporter/Exportrpc"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

func AddDE(created_by string, reason string, goods_list []byte, goods_list_amount []byte) (ok bool, err error) {
	conn, err := grpc.Dial("localhost:10004", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to exporter grpc connection fail %v", err)
		return false, err
	}
	defer conn.Close()

	c := Exporter.NewExporterClient(conn)
	r, err := c.AddDE(context.Background(), &Exporter.AddDEReq{CreatedBy: created_by, Reason: reason, GoodList: goods_list, GoodListAmount: goods_list_amount, TotalMoney: 0})
	if err != nil {
		log.Fatalf("Can't Add Don Export from boss user", err)
		return false, err
	}
	return r.Success, nil
}

func DelDE(id int64) (ok bool, err error) {
	conn, err := grpc.Dial("localhost:10004", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to exporter grpc connection fail %v", err)
		return false, err
	}
	defer conn.Close()

	c := Exporter.NewExporterClient(conn)
	r, err := c.DelDE(context.Background(), &Exporter.DelDEReq{Id: id})
	if err != nil {
		log.Fatalf("can't delete Don Export from boss user", err)
		return false, err
	}
	return r.Success, nil
}

func SelectAllDE() ([]byte, error) {
	conn, err := grpc.Dial("localhost:10004", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to exporter grpc connection fail %v", err)
		return nil, err
	}
	defer conn.Close()

	c := Exporter.NewExporterClient(conn)
	r, err := c.SelectAllDE(context.Background(), &Exporter.SelectAllDEReq{})
	if err != nil {
		log.Fatalf("can't select all Don Export from boss user", err)
		return nil, err
	}
	return r.Payload, nil
}

func SelectByIDDE(id int64) ([]byte, error) {
	conn, err := grpc.Dial("localhost:10004", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to exporter grpc connection fail %v", err)
		return nil, err
	}
	defer conn.Close()

	c := Exporter.NewExporterClient(conn)
	r, err := c.SelectByIDDE(context.Background(), &Exporter.SelectByIDDEReq{Id: id})
	if err != nil {
		log.Fatalf("can't select Don export by id from boss user", err)
		return nil, err
	}
	return r.Payload, nil
}

func SelectByCreatedAtDE(from_time int64, to_time int64) ([]byte, error) {
	conn, err := grpc.Dial("localhost:10004", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to exporter grpc connection fail %v", err)
		return nil, err
	}
	defer conn.Close()

	c := Exporter.NewExporterClient(conn)
	r, err := c.SelectByCreatedAtDE(context.Background(), &Exporter.SelectByCreatedAtDEReq{FromTime: from_time, ToTime: to_time})
	if err != nil {
		log.Fatalf("can't select Don export by created at from boss user", err)
		return nil, err
	}
	return r.Payload, nil
}
