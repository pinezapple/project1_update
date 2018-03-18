package main

import (
	"log"
	"net"
	rpc "project1_update/goods_reporter/Goodreportrpc"
	goodsreportServer "project1_update/goods_reporter/service/master"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
	rpc.RegisterGoodreportServer(rpcServer, &goodsreportServer.MasterServerImp{})

	lis, err := net.Listen("tcp", ":10009")
	if err != nil {
		log.Fatalf("failed to listen %s", err)
	}

	if err := rpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve %s", err)
	}
}
