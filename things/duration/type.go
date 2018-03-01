package duration

import "github.com/dpb587/go-schemaorg"

// // Quantity: Duration (use <a href="http://en.wikipedia.org/wiki/ISO_8601">ISO
// 8601 duration format</a>).
var Type = schemaorg.NewDataType("http://schema.org", "Duration")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
