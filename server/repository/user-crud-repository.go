package repository

import (
	"context"
	"database/sql"
	"fmt"
	"server/connection"
	"server/helper"
	"server/proto/proto_crud_user"
)

func (s *ServerUser) CreateUser(ctx context.Context, req *proto_crud_user.User) (*proto_crud_user.UserResponse, error) {

	hashPassword, _ := helper.HashPassword(req.Password)

	_, err := connection.DB.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
		req.Name, req.Email, hashPassword)
	if err != nil {
		return nil, err
	}
	return &proto_crud_user.UserResponse{Message: "User created successfully"}, nil
}

func (s *ServerUser) Login(ctx context.Context, req *proto_crud_user.LoginRequest) (*proto_crud_user.LoginResponse, error) {
	var user proto_crud_user.User

	if !helper.IsValidEmail(req.Email) {
		return &proto_crud_user.LoginResponse{
			Message: "Email harus menggunakan @gmail.com",
			Token:   "",
		}, nil
	}

	query := "SELECT id, email, password FROM users WHERE email = ?"
	err := connection.DB.QueryRow(query, req.Email).Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error while querying user: %v", err)
	}

	if !helper.VerifyPassword(req.Password, user.Password) {
		return &proto_crud_user.LoginResponse{
			Message: "password tidak cocok",
			Token:   "",
		}, nil
	}

	return &proto_crud_user.LoginResponse{
		Message: "berhasil login",
		Token:   "sdsds",
	}, nil
}

func (s *ServerUser) ReadUser(ctx context.Context, req *proto_crud_user.UserId) (*proto_crud_user.User, error) {
	var user proto_crud_user.User
	err := connection.DB.QueryRow("SELECT id, name, email, password FROM users WHERE id = ? ", req.Id).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *ServerUser) UpdateUser(ctx context.Context, req *proto_crud_user.User) (*proto_crud_user.UserResponse, error) {
	_, err := connection.DB.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", req.Name, req.Email, req.Id)
	if err != nil {
		return nil, err
	}
	return &proto_crud_user.UserResponse{Message: "User updated successfully"}, nil
}

func (s *ServerUser) DeleteUser(ctx context.Context, req *proto_crud_user.UserId) (*proto_crud_user.UserResponse, error) {
	_, err := connection.DB.Exec("DELETE FROM users WHERE id = ?", req.Id)
	if err != nil {
		return nil, err
	}
	return &proto_crud_user.UserResponse{Message: "User Delete successfully"}, nil
}

func (s *ServerUser) ListUsers(ctx context.Context, req *proto_crud_user.UserList) (*proto_crud_user.UserList, error) {
	rows, err := connection.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*proto_crud_user.User
	for rows.Next() {
		var user proto_crud_user.User
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return &proto_crud_user.UserList{Users: users}, nil
}
