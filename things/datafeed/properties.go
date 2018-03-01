package datafeed

import "github.com/dpb587/go-schemaorg"

var (
	// An item within in a data feed. Data feeds may have many elements.
	DataFeedElement = schemaorg.NewProperty("dataFeedElement")
)
