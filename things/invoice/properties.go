package invoice

import "github.com/dpb587/go-schemaorg"

var (
	// The service provider, service operator, or service performer; the goods
	// producer. Another party (a seller) may offer those services or goods on
	// behalf of the provider. A provider may also serve as the seller.
	Provider = schemaorg.NewProperty("provider")

	// The date that payment is due.
	PaymentDue = schemaorg.NewProperty("paymentDue")

	// The time interval used to compute the invoice.
	BillingPeriod = schemaorg.NewProperty("billingPeriod")

	// A number that confirms the given order or payment has been received.
	ConfirmationNumber = schemaorg.NewProperty("confirmationNumber")

	// The date the invoice is scheduled to be paid.
	ScheduledPaymentDate = schemaorg.NewProperty("scheduledPaymentDate")

	// Party placing the order or paying the invoice.
	Customer = schemaorg.NewProperty("customer")

	// A category for the item. Greater signs or slashes can be used to informally
	// indicate a category hierarchy.
	Category = schemaorg.NewProperty("category")

	// The status of payment; whether the invoice has been paid or not.
	PaymentStatus = schemaorg.NewProperty("paymentStatus")

	// An entity that arranges for an exchange between a buyer and a seller.  In
	// most cases a broker never acquires or releases ownership of a product or
	// service involved in an exchange.  If it is not clear whether an entity is a
	// broker, seller, or buyer, the latter two terms are preferred.
	Broker = schemaorg.NewProperty("broker")

	// The Order(s) related to this Invoice. One or more Orders may be combined into
	// a single Invoice.
	ReferencesOrder = schemaorg.NewProperty("referencesOrder")

	// The name of the credit card or other method of payment for the order.
	PaymentMethod = schemaorg.NewProperty("paymentMethod")

	// An identifier for the method of payment used (e.g. the last 4 digits of the
	// credit card).
	PaymentMethodId = schemaorg.NewProperty("paymentMethodId")

	// The minimum payment required at this time.
	MinimumPaymentDue = schemaorg.NewProperty("minimumPaymentDue")

	// The total amount due.
	TotalPaymentDue = schemaorg.NewProperty("totalPaymentDue")

	// The identifier for the account the payment will be applied to.
	AccountId = schemaorg.NewProperty("accountId")

	// The date that payment is due.
	PaymentDueDate = schemaorg.NewProperty("paymentDueDate")
)
