package mobileapplication

import "github.com/dpb587/go-schemaorg"

var (
	// Specifies specific carrier(s) requirements for the application (e.g. an
	// application may only work on a specific carrier network).
	CarrierRequirements = schemaorg.NewProperty("carrierRequirements")
)
