package spender

import (
	"log"
	Spender "project1_update/boss_user/microservice/spender/Spenderrpc"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

func AddDSp(created_by string, reason string, total_money int64) (ok bool, err error) {
	conn, err := grpc.Dial("localhost:10007", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to spender grpc connection fail %v", err)
		return false, err
	}
	defer conn.Close()

	c := Spender.NewSpenderClient(conn)
	r, err := c.AddDSp(context.Background(), &Spender.AddDSpReq{CreatedBy: created_by, Reason: reason, TotalMoney: total_money})
	if err != nil {
		log.Fatalf("Can't Add Don Spender from boss user", err)
		return false, err
	}
	return r.Success, nil
}

func DelDSp(id int64) (ok bool, err error) {
	conn, err := grpc.Dial("localhost:10007", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to spender grpc connection fail %v", err)
		return false, err
	}
	defer conn.Close()

	c := Spender.NewSpenderClient(conn)
	r, err := c.DelDSp(context.Background(), &Spender.DelDSpReq{Id: id})
	if err != nil {
		log.Fatalf("can't delete Don Spend from boss user", err)
		return false, err
	}
	return r.Success, nil
}

func SelectAllDSp() ([]byte, error) {
	conn, err := grpc.Dial("localhost:10007", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to spender grpc connection fail %v", err)
		return nil, err
	}
	defer conn.Close()

	c := Spender.NewSpenderClient(conn)
	r, err := c.SelectAllDSp(context.Background(), &Spender.SelectAllDSpReq{})
	if err != nil {
		log.Fatalf("can't select all Don Spend from boss user", err)
		return nil, err
	}
	return r.Payload, nil
}

func SelectByIDDSp(id int64) ([]byte, error) {
	conn, err := grpc.Dial("localhost:10007", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to spender grpc connection fail %v", err)
		return nil, err
	}
	defer conn.Close()

	c := Spender.NewSpenderClient(conn)
	r, err := c.SelectByIDDSp(context.Background(), &Spender.SelectByIDDSpReq{Id: id})
	if err != nil {
		log.Fatalf("can't select Don Spend by id from boss user", err)
		return nil, err
	}
	return r.Payload, nil
}

func SelectByCreatedAtDSp(from_time int64, to_time int64) ([]byte, error) {
	conn, err := grpc.Dial("localhost:10007", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("boss user to spender grpc connection fail %v", err)
		return nil, err
	}
	defer conn.Close()

	c := Spender.NewSpenderClient(conn)
	r, err := c.SelectByCreatedAtDSp(context.Background(), &Spender.SelectByCreatedAtDSpReq{FromTime: from_time, ToTime: to_time})
	if err != nil {
		log.Fatalf("can't select Don spend by created at from boss user", err)
		return nil, err
	}
	return r.Payload, nil
}
