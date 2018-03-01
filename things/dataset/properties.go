package dataset

import "github.com/dpb587/go-schemaorg"

var (
	// A data catalog which contains this dataset.
	IncludedInDataCatalog = schemaorg.NewProperty("includedInDataCatalog")

	// The range of temporal applicability of a dataset, e.g. for a 2011 census
	// dataset, the year 2011 (in ISO 8601 time interval format).
	DatasetTimeInterval = schemaorg.NewProperty("datasetTimeInterval")

	// A data catalog which contains this dataset.
	Catalog = schemaorg.NewProperty("catalog")

	// The International Standard Serial Number (ISSN) that identifies this serial
	// publication. You can repeat this property to identify different formats of,
	// or the linking ISSN (ISSN-L) for, this serial publication.
	Issn = schemaorg.NewProperty("issn")

	// The range of spatial applicability of a dataset, e.g. for a dataset of New
	// York weather, the state of New York.
	Spatial = schemaorg.NewProperty("spatial")

	// The range of temporal applicability of a dataset, e.g. for a 2011 census
	// dataset, the year 2011 (in ISO 8601 time interval format).
	Temporal = schemaorg.NewProperty("temporal")

	// A data catalog which contains this dataset (this property was previously
	// 'catalog', preferred name is now 'includedInDataCatalog').
	IncludedDataCatalog = schemaorg.NewProperty("includedDataCatalog")

	// A downloadable form of this dataset, at a specific location, in a specific
	// format.
	Distribution = schemaorg.NewProperty("distribution")
)
