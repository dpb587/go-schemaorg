package cookaction

import "github.com/dpb587/go-schemaorg"

var (
	// A sub property of instrument. The recipe/instructions used to perform the
	// action.
	Recipe = schemaorg.NewProperty("recipe")

	// A sub property of location. The specific food event where the action
	// occurred.
	FoodEvent = schemaorg.NewProperty("foodEvent")

	// A sub property of location. The specific food establishment where the action
	// occurred.
	FoodEstablishment = schemaorg.NewProperty("foodEstablishment")
)
