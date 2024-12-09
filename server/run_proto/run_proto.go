package runproto

import (
	"log"
	"net"
	"server/proto/proto_crud_product"
	"server/proto/proto_crud_user"
	"server/repository"

	"google.golang.org/grpc"
)

func RunnerProtos() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	proto_crud_user.RegisterCRUDServiceServer(grpcServer, &repository.ServerUser{})
	proto_crud_product.RegisterCRUDServiceServer(grpcServer, &repository.ServerProduct{})

	log.Println("gRPC server running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	// server := repository.NewServer()
	// // server := repository.Server{}

	// lis, err := net.Listen("tcp", ":50051")
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }

	// s := grpc.NewServer()
	// crudpb.RegisterCRUDServiceServer(s, server)
	// log.Println("Server is running on port 50051")
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }
}
