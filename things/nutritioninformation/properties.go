package nutritioninformation

import "github.com/dpb587/go-schemaorg"

var (
	// The number of grams of saturated fat.
	SaturatedFatContent = schemaorg.NewProperty("saturatedFatContent")

	// The number of grams of fat.
	FatContent = schemaorg.NewProperty("fatContent")

	// The number of grams of unsaturated fat.
	UnsaturatedFatContent = schemaorg.NewProperty("unsaturatedFatContent")

	// The number of grams of sugar.
	SugarContent = schemaorg.NewProperty("sugarContent")

	// The number of milligrams of cholesterol.
	CholesterolContent = schemaorg.NewProperty("cholesterolContent")

	// The number of grams of carbohydrates.
	CarbohydrateContent = schemaorg.NewProperty("carbohydrateContent")

	// The number of grams of protein.
	ProteinContent = schemaorg.NewProperty("proteinContent")

	// The number of milligrams of sodium.
	SodiumContent = schemaorg.NewProperty("sodiumContent")

	// The number of grams of trans fat.
	TransFatContent = schemaorg.NewProperty("transFatContent")

	// The number of grams of fiber.
	FiberContent = schemaorg.NewProperty("fiberContent")

	// The number of calories.
	Calories = schemaorg.NewProperty("calories")

	// The serving size, in terms of the number of volume or mass.
	ServingSize = schemaorg.NewProperty("servingSize")
)
