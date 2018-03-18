package customerDB

import (
	"log"
	cus "project1_update/orderer/microservice/services/CustomerDBrpc"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func AddToCusBalance(cusid int64, total int64) (ok bool, err error) {
	conn, err := grpc.Dial("localhost:10000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("orderer conn to customerDB err %v", err)
		return false, err
	}
	defer conn.Close()
	c := cus.NewCustomerClient(conn)

	_, err1 := c.AddToCusBalance(context.Background(), &cus.AddToBalanceReq{Id: cusid, Amount: total})
	if err1 != nil {
		log.Fatalf("can't add to cus balance %v", err1)
		return false, err1
	}
	return true, err
}

func ResetBalance(cusid int64, total int64) (ok bool, err error) {
	conn, err := grpc.Dial("localhost:10000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("orderer conn to customerDB err %v", err)
		return false, err
	}
	defer conn.Close()
	c := cus.NewCustomerClient(conn)

	_, err1 := c.AddToCusBalance(context.Background(), &cus.AddToBalanceReq{Id: cusid, Amount: -total})
	if err1 != nil {
		log.Fatalf("can't add to cus balance %v", err1)
		return false, err1
	}
	return true, err
}
