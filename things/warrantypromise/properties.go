package warrantypromise

import "github.com/dpb587/go-schemaorg"

var (
	// The duration of the warranty promise. Common unitCode values are ANN for
	// year, MON for months, or DAY for days.
	DurationOfWarranty = schemaorg.NewProperty("durationOfWarranty")

	// The scope of the warranty promise.
	WarrantyScope = schemaorg.NewProperty("warrantyScope")
)
