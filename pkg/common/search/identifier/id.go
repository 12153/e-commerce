package identifier

import "github.com/12153/e-commerce/pkg/common/search"

type ID string

// IsValid implements search.Identifier
func (id ID) IsValid() error {
	if id == "" {
		return search.IdentifierNotValidErr
	}
	return nil
}

// ToFilter implements search.Identifier
func (id ID) ToFilter() map[string]interface{} {
	return map[string]interface{}{
		"id": id,
	}
}

// ToJSON implements search.Identifier
func (id ID) ToJSON() ([]byte, error) {
	panic("unimplemented")
}

// Type implements search.Identifier
func (id ID) Type() search.Type {
	return search.IDType
}
