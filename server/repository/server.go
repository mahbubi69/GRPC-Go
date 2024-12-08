package repository

import crudpb "server/proto/proto_crud"

type Server struct {
	crudpb.UnimplementedCRUDServiceServer
	items map[string]*crudpb.Item
}

// type ServerUser struct {
// 	crudpb.UnimplementedCRUDServiceServer
// }

func NewServer() *Server {
	return &Server{
		items: make(map[string]*crudpb.Item),
	}
}
