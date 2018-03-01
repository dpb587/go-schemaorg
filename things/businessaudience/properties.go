package businessaudience

import "github.com/dpb587/go-schemaorg"

var (
	// The age of the business.
	YearsInOperation = schemaorg.NewProperty("yearsInOperation")

	// The size of the business in annual revenue.
	YearlyRevenue = schemaorg.NewProperty("yearlyRevenue")

	// The number of employees in an organization e.g. business.
	NumberOfEmployees = schemaorg.NewProperty("numberOfEmployees")
)
