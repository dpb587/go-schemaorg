package schemaorg

import "encoding/json"

type Type interface {
	json.Marshaler

	Context() string
	Type() string
}
