package game

import "github.com/dpb587/go-schemaorg"

// // The Game type represents things which are games. These are typically
// rule-governed recreational activities, e.g. role-playing games in which
// players assume the role of characters in a fictional setting.
var Type = schemaorg.NewDataType("http://schema.org", "Game")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
