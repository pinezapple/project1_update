package main

import (
	"log"
	"net"
	rpc "project1_update/spender/Spenderrpc"
	spenderServer "project1_update/spender/service/master"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
	rpc.RegisterSpenderServer(rpcServer, &spenderServer.MasterServerImp{})

	lis, err := net.Listen("tcp", ":10007")
	if err != nil {
		log.Fatalf("failed to listen %s", err)
	}

	if err := rpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve %s", err)
	}
}
