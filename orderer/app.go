package main

import (
	"log"
	"net"
	rpc "project1_update/orderer/Ordererrpc"
	ordererServer "project1_update/orderer/service/master"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
	rpc.RegisterOrdererServer(rpcServer, &ordererServer.MasterServerImp{})

	lis, err := net.Listen("tcp", ":10006")
	if err != nil {
		log.Fatalf("failed to listen %s", err)
	}

	if err := rpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve %s", err)
	}
}
