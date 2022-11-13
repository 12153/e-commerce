package search

type Identifier interface {
	Type() Type
	IsValid() error
	ToFilter() map[string]interface{}
	ToJSON() ([]byte, error)
}

type Type string

func (t Type) String() string {
	return string(t)
}

const (
	IDType       Type = "ID"
	CriteriaType Type = "Criteria"
)
