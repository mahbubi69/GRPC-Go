package repository

import (
	"context"
	"database/sql"
	"server/connection"
	"server/proto/proto_crud_user"
)

func (s *ServerUser) CreateUser(ctx context.Context, req *proto_crud_user.User) (*proto_crud_user.UserResponse, error) {
	_, err := connection.DB.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
		req.Name, req.Name, req.Email)
	if err != nil {
		return nil, err
	}
	return &proto_crud_user.UserResponse{Message: "User created successfully"}, nil
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

func (s *ServerUser) DeleteUser(ctx context.Context, req *proto_crud_user.UserId) (*proto_crud_user.UserId, error) {
	_, err := connection.DB.Exec("DELETE FROM users WHERE id = ?", req.Id)
	if err != nil {
		return nil, err
	}
	return &proto_crud_user.UserId{}, nil
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
