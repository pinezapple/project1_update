package main

import (
	"log"
	"net"
	rpc "project1_update/importer/Importerrpc"
	importerServer "project1_update/importer/service/master"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
	rpc.RegisterImporterServer(rpcServer, &importerServer.MasterServerImp{})

	lis, err := net.Listen("tcp", ":10005")
	if err != nil {
		log.Fatalf("failed to listen %s", err)
	}

	if err := rpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve %s", err)
	}
}
