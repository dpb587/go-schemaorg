package typeandquantitynode

import "github.com/dpb587/go-schemaorg"

var (
	// The unit of measurement given using the UN/CEFACT Common Code (3 characters)
	// or a URL. Other codes than the UN/CEFACT Common Code may be used with a
	// prefix followed by a colon.
	UnitCode = schemaorg.NewProperty("unitCode")

	// The product that this structured value is referring to.
	TypeOfGood = schemaorg.NewProperty("typeOfGood")

	// The quantity of the goods included in the offer.
	AmountOfThisGood = schemaorg.NewProperty("amountOfThisGood")

	// The business function (e.g. sell, lease, repair, dispose) of the offer or
	// component of a bundle (TypeAndQuantityNode). The default is
	// http://purl.org/goodrelations/v1#Sell.
	BusinessFunction = schemaorg.NewProperty("businessFunction")

	// A string or text indicating the unit of measurement. Useful if you cannot
	// provide a standard unit code for
	// <a href='unitCode'>unitCode</a>.
	UnitText = schemaorg.NewProperty("unitText")
)
