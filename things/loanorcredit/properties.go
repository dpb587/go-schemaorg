package loanorcredit

import "github.com/dpb587/go-schemaorg"

var (
	// The duration of the loan or credit agreement.
	LoanTerm = schemaorg.NewProperty("loanTerm")

	// The amount of money.
	Amount = schemaorg.NewProperty("amount")

	// Assets required to secure loan or credit repayments. It may take form of
	// third party pledge, goods, financial instruments (cash, securities, etc.)
	RequiredCollateral = schemaorg.NewProperty("requiredCollateral")
)
