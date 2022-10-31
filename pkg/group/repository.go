package group

import (
	"context"

	"github.com/gofrs/uuid"
)

type Repository interface {
	CreateGroup(context.Context, CreateGroupRequest) (CreateGroupResponse, error)
	RetrieveGroup(context.Context, RetrieveGroupRequest) (RetrieveGroupResponse, error)
	DeleteGroup(context.Context, DeleteGroupRequest) (DeleteGroupResponse, error)
	SearchGroup(context.Context, SearchGroupRequest) (SearchGroupResponse, error)
}

type CreateGroupRequest struct {
}
type CreateGroupResponse struct {
	Group Group `json:"group"`
}

type RetrieveGroupRequest struct {
}
type RetrieveGroupResponse struct {
	Group Group `json:"group"`
}

type DeleteGroupRequest struct {
	ID uuid.UUID
}
type DeleteGroupResponse struct {
}

type SearchGroupRequest struct {
}
type SearchGroupResponse struct {
}
