package financialproduct

import "github.com/dpb587/go-schemaorg"

var (
	// The annual rate that is charged for borrowing (or made by investing),
	// expressed as a single percentage number that represents the actual yearly
	// cost of funds over the term of a loan. This includes any fees or additional
	// costs associated with the transaction.
	AnnualPercentageRate = schemaorg.NewProperty("annualPercentageRate")

	// The interest rate, charged or paid, applicable to the financial product.
	// Note: This is different from the calculated annualPercentageRate.
	InterestRate = schemaorg.NewProperty("interestRate")

	// Description of fees, commissions, and other terms applied either to a class
	// of financial product, or by a financial service organization.
	FeesAndCommissionsSpecification = schemaorg.NewProperty("feesAndCommissionsSpecification")
)
