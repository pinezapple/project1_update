package main

import (
	"log"
	"net"
	rpc "project1_update/goodsDB/GoodsDBrpc"
	goodsServer "project1_update/goodsDB/service/master"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
	rpc.RegisterGoodsServer(rpcServer, &goodsServer.MasterServerImp{})

	lis, err := net.Listen("tcp", ":10001")
	if err != nil {
		log.Fatalf("failed to listen %s", err)
	}

	if err := rpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve %s", err)
	}
}
