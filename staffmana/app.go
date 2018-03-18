package main

import (
	"log"
	"net"
	rpc "project1_update/staffmana/Staffmanarpc"
	staffmanaServer "project1_update/staffmana/service/master"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
	rpc.RegisterStaffmanaServer(rpcServer, &staffmanaServer.MasterServerImp{})

	lis, err := net.Listen("tcp", ":10010")
	if err != nil {
		log.Fatalf("failed to listen %s", err)
	}

	if err := rpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve %s", err)
	}
}
