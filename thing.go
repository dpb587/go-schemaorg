package schemaorg

import (
	"encoding/json"
)

type Thing struct {
	ourType    Type
	properties map[string][]Type
}

var _ Type = &Thing{}

func NewThing(ourType Type) *Thing {
	return &Thing{
		ourType:    ourType,
		properties: map[string][]Type{},
	}
}

func (d *Thing) Context() string {
	return d.ourType.Context()
}

func (d *Thing) Type() string {
	return d.ourType.Type()
}

func (d *Thing) Add(property Property, data Type) *Thing {
	d.properties[property.Name()] = append(d.properties[property.Name()], data)

	return d
}

func (d *Thing) MarshalJSON() ([]byte, error) {
	result := map[string]interface{}{}

	result["@type"] = d.ourType.Type()
	result["@context"] = d.ourType.Context()

	id, idd := d.properties["identifier"]
	if idd {
		result["@id"] = id[0]
	}

	for propertyName, propertyValues := range d.properties {
		if len(propertyValues) == 0 {
			continue
		} else if len(propertyValues) == 1 {
			result[propertyName] = propertyValues[0]
		} else {
			result[propertyName] = propertyValues
		}
	}

	return json.Marshal(result)
}
