package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"github.com/thanhauco/nexus/common/proto"
)

type server struct {
	proto.UnimplementedCatalogServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterCatalogServiceServer(s, &server{})
	log.Printf("Catalog Service listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
