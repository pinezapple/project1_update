package main

import (
	"log"
	"net"
	rpc "project1_update/exporter/Exportrpc"
	exporterServer "project1_update/exporter/service/master"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
	rpc.RegisterExporterServer(rpcServer, &exporterServer.MasterServerImp{})

	lis, err := net.Listen("tcp", ":10004")
	if err != nil {
		log.Fatalf("failed to listen %s", err)
	}

	if err := rpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve %s", err)
	}
}
