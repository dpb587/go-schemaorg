package datedmoneyspecification

import "github.com/dpb587/go-schemaorg"

var (
	// The end date and time of the item (in <a
	// href="http://en.wikipedia.org/wiki/ISO_8601">ISO 8601 date format</a>).
	EndDate = schemaorg.NewProperty("endDate")

	// The currency in which the monetary amount is expressed (in 3-letter <a
	// href="http://en.wikipedia.org/wiki/ISO_4217">ISO 4217</a> format).
	Currency = schemaorg.NewProperty("currency")

	// The amount of money.
	Amount = schemaorg.NewProperty("amount")

	// The start date and time of the item (in <a
	// href="http://en.wikipedia.org/wiki/ISO_8601">ISO 8601 date format</a>).
	StartDate = schemaorg.NewProperty("startDate")
)
