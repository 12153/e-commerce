package company

import (
	"context"

	"github.com/gofrs/uuid"
)

type Repository interface {
	CreateCompany(context.Context, CreateCompanyRequest) (CreateCompanyResponse, error)
	RetrieveCompany(context.Context, RetrieveCompanyRequest) (RetrieveCompanyResponse, error)
	DeleteCompany(context.Context, DeleteCompanyRequest) (DeleteCompanyResponse, error)
	SearchCompany(context.Context, SearchCompanyRequest) (SearchCompanyResponse, error)
}

type CreateCompanyRequest struct {
}
type CreateCompanyResponse struct {
	Company Company `json:"Company"`
}

type RetrieveCompanyRequest struct {
}
type RetrieveCompanyResponse struct {
	Company Company `json:"Company"`
}

type DeleteCompanyRequest struct {
	ID uuid.UUID
}
type DeleteCompanyResponse struct {
}

type SearchCompanyRequest struct {
}
type SearchCompanyResponse struct {
}
