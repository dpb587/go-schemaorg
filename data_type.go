package schemaorg

import "encoding/json"

type DataType struct {
	context_ string
	type_    string
}

func NewDataType(context_, type_ string) DataType {
	return DataType{
		context_: context_,
		type_:    type_,
	}
}

func (v DataType) Context() string {
	return v.context_
}

func (v DataType) Type() string {
	return v.type_
}

func (v DataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"@context": v.context_,
		"@type":    v.type_,
	})
}
