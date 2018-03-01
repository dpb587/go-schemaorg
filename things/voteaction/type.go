package voteaction

import "github.com/dpb587/go-schemaorg"

// // The act of expressing a preference from a fixed/finite/structured set of
// choices/options.
var Type = schemaorg.NewDataType("http://schema.org", "VoteAction")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
