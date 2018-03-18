package importer

import (
	"log"
	Importer "project1_update/boss_user/microservice/importer/Importerrpc"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

func AddDI(created_by string, goods_list []byte, goods_list_amount []byte) (ok bool, err error) {
	conn, err := grpc.Dial("localhost:10005", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to importer grpc connection fail %v", err)
		return false, err
	}
	defer conn.Close()

	c := Importer.NewImporterClient(conn)
	r, err := c.AddDI(context.Background(), &Importer.AddDIReq{CreatedBy: created_by, GoodList: goods_list, GoodListAmount: goods_list_amount, TotalMoney: 0})
	if err != nil {
		log.Fatalf("Can't Add Don Import from boss user", err)
		return false, err
	}
	return r.Success, nil
}

func DelDI(id int64) (ok bool, err error) {
	conn, err := grpc.Dial("localhost:10005", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to importer grpc connection fail %v", err)
		return false, err
	}
	defer conn.Close()

	c := Importer.NewImporterClient(conn)
	r, err := c.DelDI(context.Background(), &Importer.DelDIReq{Id: id})
	if err != nil {
		log.Fatalf("can't delete Don Import from boss user", err)
		return false, err
	}
	return r.Success, nil
}

func SelectAllDI() ([]byte, error) {
	conn, err := grpc.Dial("localhost:10005", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to importer grpc connection fail %v", err)
		return nil, err
	}
	defer conn.Close()

	c := Importer.NewImporterClient(conn)
	r, err := c.SelectAllDI(context.Background(), &Importer.SelectAllDIReq{})
	if err != nil {
		log.Fatalf("can't select all Don Import from boss user", err)
		return nil, err
	}
	return r.Payload, nil
}

func SelectByIDDI(id int64) ([]byte, error) {
	conn, err := grpc.Dial("localhost:10005", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to importer grpc connection fail %v", err)
		return nil, err
	}
	defer conn.Close()

	c := Importer.NewImporterClient(conn)
	r, err := c.SelectByIDDI(context.Background(), &Importer.SelectByIDDIReq{Id: id})
	if err != nil {
		log.Fatalf("can't select Don import by id from boss user", err)
		return nil, err
	}
	return r.Payload, nil
}

func SelectByCreatedAtDI(from_time int64, to_time int64) ([]byte, error) {
	conn, err := grpc.Dial("localhost:10005", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to importer grpc connection fail %v", err)
		return nil, err
	}
	defer conn.Close()

	c := Importer.NewImporterClient(conn)
	r, err := c.SelectByCreatedAtDI(context.Background(), &Importer.SelectByCreatedAtDIReq{FromTime: from_time, ToTime: to_time})
	if err != nil {
		log.Fatalf("can't select Don import by created at from boss user", err)
		return nil, err
	}
	return r.Payload, nil
}
