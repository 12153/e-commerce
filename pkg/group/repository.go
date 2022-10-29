package group

import "github.com/gofrs/uuid"

type Repository interface {
	CreateGroup(CreateGroupRequest) (CreateGroupResponse, error)
	RetrieveGroup(RetrieveGroupRequest) (RetrieveGroupResponse, error)
	DeleteGroup(DeleteGroupRequest) (DeleteGroupResponse, error)
	SearchGroup(SearchGroupRequest) (SearchGroupResponse, error)
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
