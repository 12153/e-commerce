package user

import "github.com/gofrs/uuid"

type User struct {
	ID      uuid.UUID
	OwnerID uuid.UUID
	Name    string
	Email   string
}
