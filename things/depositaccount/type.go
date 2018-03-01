package depositaccount

import "github.com/dpb587/go-schemaorg"

// // A type of Bank Account with a main purpose of depositing funds to gain
// interest or other benefits.
var Type = schemaorg.NewDataType("http://schema.org", "DepositAccount")

func New() *schemaorg.Thing {
	return schemaorg.NewThing(Type)
}
