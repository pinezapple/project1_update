package cusmana

import (
	"log"

	Cusmana "project1_update/boss_user/microservice/cusmana/Cusmanarpc"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

func AddCus(name string, phonenum string, balance int64) (ok bool, err error) {

	conn, err := grpc.Dial("localhost:10008", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to cusmana grpc connection fail %v", err)
		return false, err
	}
	defer conn.Close()
	c := Cusmana.NewCusmanaClient(conn)

	r, err := c.AddCus(context.Background(), &Cusmana.AddCusReq{Name: name, Phonenum: phonenum, Balance: balance})

	if err != nil {
		log.Fatalf("Can't addcus from boss user %v", err)
		return false, err
	}

	return r.Success, nil
}

func DelCus(id int64) (ok bool, err error) {

	conn, err := grpc.Dial("localhost:10008", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to cusmana grpc connection fail %v", err)
		return false, err
	}
	defer conn.Close()
	c := Cusmana.NewCusmanaClient(conn)

	r, err := c.DelCus(context.Background(), &Cusmana.DelCusReq{Id: id})

	if err != nil {
		log.Fatalf("Can't delcus from boss user %v", err)
		return false, err
	}

	return r.Success, nil
}

func UpdateCus(id int64, name string, phonenum string, balance int64, level int64) (ok bool, err error) {

	conn, err := grpc.Dial("localhost:10008", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to cusmana grpc connection fail %v", err)
		return false, err
	}
	defer conn.Close()
	c := Cusmana.NewCusmanaClient(conn)

	r, err := c.UpdateCus(context.Background(), &Cusmana.UpdateCusReq{Id: id, Name: name, Phonenum: phonenum, Balance: balance, Level: level})

	if err != nil {
		log.Fatalf("Can't updatecus from boss user %v", err)
		return false, err
	}

	return r.Success, nil
}

func SelectAllCus() (mar []byte, err error) {
	conn, err := grpc.Dial("localhost:10008", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to cusmana grpc connection fail %v", err)
		return nil, err
	}
	defer conn.Close()
	c := Cusmana.NewCusmanaClient(conn)

	r, err := c.SelectAllCus(context.Background(), &Cusmana.SelectAllCusReq{})

	if err != nil {
		log.Fatalf("Can't select all cus from boss user", err)
		return nil, err
	}

	return r.Payload, err
}

func SelectCusById(id int64) (mar []byte, err error) {
	conn, err := grpc.Dial("localhost:10008", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to cusmana grpc connection fail %v", err)
		return nil, err
	}
	defer conn.Close()
	c := Cusmana.NewCusmanaClient(conn)

	r, err1 := c.SelectCusById(context.Background(), &Cusmana.SelectCusByIdReq{Id: id})
	if err1 != nil {
		log.Fatalf("can't select customers by id from boss user %v", err1)
		return nil, err1
	}

	return r.Payload, err
}

func SelectCusByName(name string) (mar []byte, err error) {
	conn, err := grpc.Dial("localhost:10008", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to cusmana grpc connection fail %v", err)
		return nil, err
	}
	defer conn.Close()
	c := Cusmana.NewCusmanaClient(conn)

	r, err1 := c.SelectCusByName(context.Background(), &Cusmana.SelectCusByNameReq{Name: name})
	if err1 != nil {
		log.Fatalf("can't select customers by name from boss user %v", err1)
		return nil, err1
	}

	return r.Payload, err
}

func SelectCusByPhoneNum(phonenum string) (mar []byte, err error) {
	conn, err := grpc.Dial("localhost:10008", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to cusmana grpc connection fail %v", err)
		return nil, err
	}
	defer conn.Close()
	c := Cusmana.NewCusmanaClient(conn)

	r, err1 := c.SelectCusByPhoneNum(context.Background(), &Cusmana.SelectCusByPhoneNumReq{Phonenum: phonenum})
	if err1 != nil {
		log.Fatalf("can't select customers by phonenum from boss user %v", err1)
		return nil, err1
	}

	return r.Payload, err
}
