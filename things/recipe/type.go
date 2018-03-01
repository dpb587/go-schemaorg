package recipe

import "github.com/dpb587/go-schemaorg"

// // A recipe. For dietary restrictions covered by the recipe, a few common
// restrictions are enumerated via <a class="localLink"
// href="http://schema.org/suitableForDiet">suitableForDiet</a>. The <a
// class="localLink" href="http://schema.org/keywords">keywords</a> property can
// also be used to add more detail.
var Type = schemaorg.NewDataType("http://schema.org", "Recipe")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
