package menusection

import "github.com/dpb587/go-schemaorg"

// // A sub-grouping of food or drink items in a menu. E.g. courses (such as
// 'Dinner', 'Breakfast', etc.), specific type of dishes (such as 'Meat',
// 'Vegan', 'Drinks', etc.), or some other classification made by the menu
// provider.
var Type = schemaorg.NewDataType("http://schema.org", "MenuSection")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
