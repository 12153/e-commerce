package company

import (
	"time"

	"github.com/gofrs/uuid"
)

type Company struct {
	ID               uuid.UUID
	OwnerID          uuid.UUID
	Name             string
	Description      string
	RegistrationDate time.Time
}
