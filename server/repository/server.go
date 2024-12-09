package repository

import (
	"server/proto/proto_crud_product"
	crudpb "server/proto/proto_crud_test"
	"server/proto/proto_crud_user"
)

type Server struct {
	crudpb.UnimplementedCRUDServiceServer
	items map[string]*crudpb.Item
}

func NewServer() *Server {
	return &Server{
		items: make(map[string]*crudpb.Item),
	}
}

type ServerUser struct {
	proto_crud_user.UnimplementedCRUDServiceServer
}

type ServerProduct struct {
	proto_crud_product.UnimplementedCRUDServiceServer
}
