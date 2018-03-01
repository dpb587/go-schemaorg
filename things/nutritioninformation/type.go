package nutritioninformation

import "github.com/dpb587/go-schemaorg"

// // Nutritional information about the recipe.
var Type = schemaorg.NewDataType("http://schema.org", "NutritionInformation")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
