package repository

import (
	"context"
	crudpb "server/proto/proto_crud_test"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func (s *Server) CreateItem(ctx context.Context, req *crudpb.CreateItemRequest) (*crudpb.CreateItemResponse, error) {
	item := req.GetItem()
	s.items[item.Id] = item
	return &crudpb.CreateItemResponse{Item: item}, nil
}

func (s *Server) ReadItem(ctx context.Context, req *crudpb.ReadItemRequest) (*crudpb.ReadItemResponse, error) {
	item, exists := s.items[req.GetId()]
	if !exists {
		return nil, grpc.Errorf(codes.NotFound, "item not found")
	}
	return &crudpb.ReadItemResponse{Item: item}, nil
}

func (s *Server) UpdateItem(ctx context.Context, req *crudpb.UpdateItemRequest) (*crudpb.UpdateItemResponse, error) {
	item := req.GetItem()
	s.items[item.Id] = item
	return &crudpb.UpdateItemResponse{Item: item}, nil
}

func (s *Server) DeleteItem(ctx context.Context, req *crudpb.DeleteItemRequest) (*crudpb.DeleteItemResponse, error) {
	delete(s.items, req.GetId())
	return &crudpb.DeleteItemResponse{Id: req.GetId()}, nil
}
