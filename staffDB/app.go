package main

import (
	"log"
	"net"
	rpc "project1_update/staffDB/StaffDBrpc"
	staffServer "project1_update/staffDB/service/master"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
	rpc.RegisterStaffServer(rpcServer, &staffServer.MasterServerImp{})

	lis, err := net.Listen("tcp", ":10002")
	if err != nil {
		log.Fatalf("failed to listen %s", err)
	}

	if err := rpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve %s", err)
	}
}
