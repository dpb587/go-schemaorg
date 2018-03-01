package integer

import (
	"encoding/json"
	"github.com/dpb587/go-schemaorg"
)

type Type struct {
	value int64
}

var _ schemaorg.Type = Type{}

func New(value int64) Type {
	return Type{value: value}
}

func (Type) Context() string {
	return "http://schema.org"
}

func (Type) Type() string {
	return "Integer"
}

func (t Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.value)
}
