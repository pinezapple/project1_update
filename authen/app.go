package main

import (
	"log"
	"net"
	rpc "project1_update/authen/Authenrpc"
	authenServer "project1_update/authen/service/master"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
	rpc.RegisterAuthenServer(rpcServer, &authenServer.MasterServerImp{})

	lis, err := net.Listen("tcp", ":10003")
	if err != nil {
		log.Fatalf("failed to listen %s", err)
	}

	if err := rpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve %s", err)
	}
}
