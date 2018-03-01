package time

import (
	"encoding/json"
	"github.com/dpb587/go-schemaorg"
	"time"
)

type Type struct {
	value time.Time
}

var _ schemaorg.Type = Type{}

func New(value time.Time) Type {
	return Type{value: value}
}

func (Type) Context() string {
	return "http://schema.org"
}

func (Type) Type() string {
	return "Time"
}

func (t Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.value.Format("T15:04:05Z07:00"))
}
