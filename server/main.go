package main

import (
	"fmt"
	pb "github.com/MeibisuX673/grpc/proto/infoDirectory"
	"github.com/MeibisuX673/grpc/server/services/directoryInfo"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 5300))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterInfoDirectoryServer(grpcServer, &directoryInfo.DirectoryInfo{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
