package foodevent

import "github.com/dpb587/go-schemaorg"

// // Event type: Food event.
var Type = schemaorg.NewDataType("http://schema.org", "FoodEvent")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
