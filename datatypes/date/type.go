package date

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
	return "Date"
}

func (t Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.value.Format("2006-01-02"))
}
