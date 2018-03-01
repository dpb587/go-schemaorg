package float

import (
	"encoding/json"
	"github.com/dpb587/go-schemaorg"
)

type Type struct {
	value float64
}

var _ schemaorg.Type = Type{}

func New(value float64) Type {
	return Type{value: value}
}

func (Type) Context() string {
	return "http://schema.org"
}

func (Type) Type() string {
	return "Float"
}

func (t Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.value)
}
