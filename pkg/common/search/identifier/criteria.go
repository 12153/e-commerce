package identifier

import (
	"encoding/json"

	"github.com/12153/e-commerce/pkg/common/search"
)

type Criterion struct {
	Key   string
	Value interface{}
}

type Criteria []Criterion

// IsValid implements search.Identifier
func (c Criteria) IsValid() error {
	if len(c) == 0 {
		return search.IdentifierNotValidErr
	}
	return nil
}

// ToFilter implements search.Identifier
func (c Criteria) ToFilter() map[string]interface{} {
	filter := make(map[string]interface{})
	for _, v := range c {
		filter[v.Key] = v.Value
	}
	return filter
}

// ToJSON implements search.Identifier
func (c Criteria) ToJSON() ([]byte, error) {
	return json.MarshalIndent(c, "", "  ")
}

// Type implements search.Identifier
func (c Criteria) Type() search.Type {
	panic("unimplemented")
}
