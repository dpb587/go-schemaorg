package computerlanguage

import "github.com/dpb587/go-schemaorg"

// // This type covers computer programming languages such as Scheme and Lisp, as
// well as other language-like computer representations. Natural languages are
// best represented with the <a class="localLink"
// href="http://schema.org/Language">Language</a> type.
var Type = schemaorg.NewDataType("http://schema.org", "ComputerLanguage")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
