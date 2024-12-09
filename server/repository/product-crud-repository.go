package repository

import (
	"context"
	"database/sql"
	"server/connection"
	"server/proto/proto_crud_product"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s *ServerProduct) CreateProduct(ctx context.Context, req *proto_crud_product.Product) (*proto_crud_product.ProductResponse, error) {

	_, err := connection.DB.Exec("INSERT INTO product (name, category, item) VALUES (?, ?, ?)", req.Name, req.Category, req.Item)

	if err != nil {
		return nil, err
	}

	grpc.SetHeader(ctx, metadata.Pairs("http-status", "200"))

	return &proto_crud_product.ProductResponse{
		Message: "Product created successfully",
		Success: true,
		Status:  "200 OK",
	}, nil
}

func (s *ServerProduct) ReadProduct(ctx context.Context, req *proto_crud_product.ProductId) (*proto_crud_product.Product, error) {

	var product proto_crud_product.Product

	err := connection.DB.QueryRow("SELECT id, name, category, item FROM product WHERE id = ? ", req.Id).Scan(&product.Id, &product.Name, &product.Category, &product.Item)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (s *ServerProduct) UpdateProduct(ctx context.Context, req *proto_crud_product.Product) (*proto_crud_product.ProductResponse, error) {
	_, err := connection.DB.Exec("UPDATE product SET name = ? , category = ?, item = ? WHERE id = ? ", req.Name, req.Category, req.Item, req.Id)

	if err != nil {
		return nil, err
	}

	return &proto_crud_product.ProductResponse{
			Message: "Update Product Succesfully",
			Success: true,
			Status:  "200"},
		nil
}

func (s *ServerProduct) DeleteProduct(ctx context.Context, req *proto_crud_product.ProductId) (*proto_crud_product.ProductResponse, error) {
	_, err := connection.DB.Exec("DELETE FROM product WHERE id = ?", req.Id)

	if err != nil {
		return nil, err
	}

	return &proto_crud_product.ProductResponse{
			Message: "Delete Product Succesfully",
			Success: true,
			Status:  "200"},
		nil

}

func (s *ServerProduct) ListProducts(ctx context.Context, req *proto_crud_product.ProductList) (*proto_crud_product.ProductList, error) {

	rows, err := connection.DB.Query("SELECT id, name, category, item FROM product")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*proto_crud_product.Product
	// var count uint32
	for rows.Next() {
		var product proto_crud_product.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Category, &product.Item)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
		// count++
	}

	count := len(products)

	return &proto_crud_product.ProductList{
			Status:   "200",
			Count:    uint32(count),
			Products: products},
		nil
}
