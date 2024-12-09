package main

import (
	"context"
	"log"
	"time"

	crudpb "server/proto/proto_crud_test"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := crudpb.NewCRUDServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// // Create an item
	item := &crudpb.Item{Id: "1", Name: "Item 1", Description: "This is item 1"}
	res, err := c.CreateItem(ctx, &crudpb.CreateItemRequest{Item: item})
	if err != nil {
		log.Fatalf("could not create item: %v", err)
	}
	log.Printf("Item created: %v", res.GetItem())

	// Read an item
	resRead, err := c.ReadItem(ctx, &crudpb.ReadItemRequest{Id: "1"})
	if err != nil {
		log.Fatalf("could not read item: %v", err)
	}
	log.Printf("Item read: %v", resRead.GetItem())

	// // Update an item
	// item.Description = "This is the updated item 1"
	// resUpdate, err := c.UpdateItem(ctx, &crudpb.UpdateItemRequest{Item: item})
	// if err != nil {
	// 	log.Fatalf("could not update item: %v", err)
	// }
	// log.Printf("Item updated: %v", resUpdate.GetItem())

	// // Delete an item
	// resDelete, err := c.DeleteItem(ctx, &crudpb.DeleteItemRequest{Id: "1"})
	// if err != nil {
	// 	log.Fatalf("could not delete item: %v", err)
	// }
	// log.Printf("Item deleted: %v", resDelete.GetId())
}
