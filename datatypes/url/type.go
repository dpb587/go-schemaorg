package url

import (
	"encoding/json"

	"github.com/dpb587/go-schemaorg"
)

type Type struct {
	value string
}

var _ schemaorg.Type = Type{}

func New(value string) Type {
	return Type{value: value}
}

func (Type) Context() string {
	return "http://schema.org"
}

func (Type) Type() string {
	return "URL"
}

func (t Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.value)
}
