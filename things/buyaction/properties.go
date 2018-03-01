package buyaction

import "github.com/dpb587/go-schemaorg"

var (
	// 'vendor' is an earlier term for 'seller'.
	Vendor = schemaorg.NewProperty("vendor")

	// An entity which offers (sells / leases / lends / loans) the services / goods.
	// A seller may also be a provider.
	Seller = schemaorg.NewProperty("seller")

	// The warranty promise(s) included in the offer.
	WarrantyPromise = schemaorg.NewProperty("warrantyPromise")
)
