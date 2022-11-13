package product

import "github.com/gofrs/uuid"

type Product struct {
	ID            uuid.UUID `json:"id"`
	OwnerID       uuid.UUID `json:"owner_id"`
	Name          string
	Price         float64
	ImagesURLList []string
}
