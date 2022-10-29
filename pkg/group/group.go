package group

import "github.com/gofrs/uuid"

type Group struct {
	ID      uuid.UUID `json:"id"`
	OwnerID uuid.UUID `json:"ownerID"`
}

const GroupTypeName = "commerce://Group"
