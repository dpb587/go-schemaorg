package menuitem

import "github.com/dpb587/go-schemaorg"

var (
	// Nutrition information about the recipe or menu item.
	Nutrition = schemaorg.NewProperty("nutrition")

	// An offer to provide this item&#x2014;for example, an offer to sell a product,
	// rent the DVD of a movie, perform a service, or give away tickets to an event.
	Offers = schemaorg.NewProperty("offers")

	// Indicates a dietary restriction or guideline for which this recipe or menu
	// item is suitable, e.g. diabetic, halal etc.
	SuitableForDiet = schemaorg.NewProperty("suitableForDiet")
)
