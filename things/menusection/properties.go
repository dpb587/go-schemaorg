package menusection

import "github.com/dpb587/go-schemaorg"

var (
	// A food or drink item contained in a menu or menu section.
	HasMenuItem = schemaorg.NewProperty("hasMenuItem")

	// A subgrouping of the menu (by dishes, course, serving time period, etc.).
	HasMenuSection = schemaorg.NewProperty("hasMenuSection")
)
