package main

import (
	"crypto/tls"
	"fmt"
	pb "github.com/MeibisuX673/grpc/server/proto/infoDirectory"
	"github.com/MeibisuX673/grpc/server/services/directoryInfo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

var (
	crtFile = "server.crt"
	keyFile = "server.key"
)

func main() {

	cert, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}
	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 5300))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterInfoDirectoryServer(grpcServer, &directoryInfo.DirectoryInfo{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
