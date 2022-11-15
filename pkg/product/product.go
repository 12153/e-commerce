package product

import "github.com/gofrs/uuid"

type Product struct {
	ID            uuid.UUID `json:"id"`
	OwnerID       uuid.UUID `json:"owner_id"`
	Name          string    `bson:"name"`
	Price         float64   `bson:"price"`
	ImagesURLList []string  `bson:"imagesurllist"`
}
