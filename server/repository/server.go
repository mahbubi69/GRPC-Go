package repository

import (
	crudpb "server/proto/proto_crud"
	"server/proto/proto_crud_user"
)

type Server struct {
	crudpb.UnimplementedCRUDServiceServer
	items map[string]*crudpb.Item
}
type ServerUser struct {
	proto_crud_user.UnimplementedCRUDServiceServer
}

func NewServer() *Server {
	return &Server{
		items: make(map[string]*crudpb.Item),
	}
}
