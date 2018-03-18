package main

import (
	"log"
	"net"
	rpc "project1_update/customerDB/CustomerDBrpc"
	cusServer "project1_update/customerDB/service/master"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
	rpc.RegisterCustomerServer(rpcServer, &cusServer.MasterServerImp{})

	lis, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatalf("failed to listen %s", err)
	}

	if err := rpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve %s", err)
	}
}
