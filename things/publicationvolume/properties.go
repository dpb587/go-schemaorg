package publicationvolume

import "github.com/dpb587/go-schemaorg"

var (
	// Any description of pages that is not separated into pageStart and pageEnd;
	// for example, "1-6, 9, 55" or "10-12, 46-49".
	Pagination = schemaorg.NewProperty("pagination")

	// The page on which the work ends; for example "138" or "xvi".
	PageEnd = schemaorg.NewProperty("pageEnd")

	// Identifies the volume of publication or multi-part work; for example, "iii"
	// or "2".
	VolumeNumber = schemaorg.NewProperty("volumeNumber")

	// The page on which the work starts; for example "135" or "xiii".
	PageStart = schemaorg.NewProperty("pageStart")
)
