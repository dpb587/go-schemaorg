package recipe

import "github.com/dpb587/go-schemaorg"

var (
	// The quantity produced by the recipe (for example, number of people served,
	// number of servings, etc).
	RecipeYield = schemaorg.NewProperty("recipeYield")

	// The category of the recipe—for example, appetizer, entree, etc.
	RecipeCategory = schemaorg.NewProperty("recipeCategory")

	// A single ingredient used in the recipe, e.g. sugar, flour or garlic.
	RecipeIngredient = schemaorg.NewProperty("recipeIngredient")

	// A step in making the recipe, in the form of a single item (document, video,
	// etc.) or an ordered list with HowToStep and/or HowToSection items.
	RecipeInstructions = schemaorg.NewProperty("recipeInstructions")

	// Nutrition information about the recipe or menu item.
	Nutrition = schemaorg.NewProperty("nutrition")

	// Indicates a dietary restriction or guideline for which this recipe or menu
	// item is suitable, e.g. diabetic, halal etc.
	SuitableForDiet = schemaorg.NewProperty("suitableForDiet")

	// The cuisine of the recipe (for example, French or Ethiopian).
	RecipeCuisine = schemaorg.NewProperty("recipeCuisine")

	// The time it takes to actually cook the dish, in <a
	// href="http://en.wikipedia.org/wiki/ISO_8601">ISO 8601 duration format</a>.
	CookTime = schemaorg.NewProperty("cookTime")

	// The method of cooking, such as Frying, Steaming, ...
	CookingMethod = schemaorg.NewProperty("cookingMethod")

	// A single ingredient used in the recipe, e.g. sugar, flour or garlic.
	Ingredients = schemaorg.NewProperty("ingredients")
)
