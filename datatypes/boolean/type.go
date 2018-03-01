package boolean

import (
	"encoding/json"
	"github.com/dpb587/go-schemaorg"
)

type Type struct {
	value bool
}

var _ schemaorg.Type = Type{}

func New(value bool) Type {
	return Type{value: value}
}

func (Type) Context() string {
	return "http://schema.org"
}

func (t Type) Type() string {
	switch t.value {
	case true:
		return "True"
	case false:
		return "False"
	}

	panic("unreachable code")
}

func (t Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.value)
}
