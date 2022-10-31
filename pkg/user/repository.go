package user

import (
	"context"

	"github.com/gofrs/uuid"
)

type Repository interface {
	CreateUser(context.Context, CreateUserRequest) (CreateUserResponse, error)
	RetrieveUser(context.Context, RetrieveUserRequest) (RetrieveUserResponse, error)
	DeleteUser(context.Context, DeleteUserRequest) (DeleteUserResponse, error)
	SearchUser(context.Context, SearchUserRequest) (SearchUserResponse, error)
}

type CreateUserRequest struct {
}
type CreateUserResponse struct {
	User User `json:"User"`
}

type RetrieveUserRequest struct {
}
type RetrieveUserResponse struct {
	User User `json:"User"`
}

type DeleteUserRequest struct {
	ID uuid.UUID
}
type DeleteUserResponse struct {
}

type SearchUserRequest struct {
}
type SearchUserResponse struct {
}
