package order

import "github.com/dpb587/go-schemaorg"

var (
	// Date order was placed.
	OrderDate = schemaorg.NewProperty("orderDate")

	// The date that payment is due.
	PaymentDue = schemaorg.NewProperty("paymentDue")

	// The identifier of the transaction.
	OrderNumber = schemaorg.NewProperty("orderNumber")

	// The URL for sending a payment.
	PaymentUrl = schemaorg.NewProperty("paymentUrl")

	// The current status of the order.
	OrderStatus = schemaorg.NewProperty("orderStatus")

	// The billing address for the order.
	BillingAddress = schemaorg.NewProperty("billingAddress")

	// The order is being paid as part of the referenced Invoice.
	PartOfInvoice = schemaorg.NewProperty("partOfInvoice")

	// A number that confirms the given order or payment has been received.
	ConfirmationNumber = schemaorg.NewProperty("confirmationNumber")

	// Party placing the order or paying the invoice.
	Customer = schemaorg.NewProperty("customer")

	// 'merchant' is an out-dated term for 'seller'.
	Merchant = schemaorg.NewProperty("merchant")

	// An entity that arranges for an exchange between a buyer and a seller.  In
	// most cases a broker never acquires or releases ownership of a product or
	// service involved in an exchange.  If it is not clear whether an entity is a
	// broker, seller, or buyer, the latter two terms are preferred.
	Broker = schemaorg.NewProperty("broker")

	// Was the offer accepted as a gift for someone other than the buyer.
	IsGift = schemaorg.NewProperty("isGift")

	// The name of the credit card or other method of payment for the order.
	PaymentMethod = schemaorg.NewProperty("paymentMethod")

	// An entity which offers (sells / leases / lends / loans) the services / goods.
	// A seller may also be a provider.
	Seller = schemaorg.NewProperty("seller")

	// An identifier for the method of payment used (e.g. the last 4 digits of the
	// credit card).
	PaymentMethodId = schemaorg.NewProperty("paymentMethodId")

	// Any discount applied (to an Order).
	Discount = schemaorg.NewProperty("discount")

	// The delivery of the parcel related to this order or order item.
	OrderDelivery = schemaorg.NewProperty("orderDelivery")

	// The item ordered.
	OrderedItem = schemaorg.NewProperty("orderedItem")

	// The offer(s) -- e.g., product, quantity and price combinations -- included in
	// the order.
	AcceptedOffer = schemaorg.NewProperty("acceptedOffer")

	// The currency (in 3-letter ISO 4217 format) of the discount.
	DiscountCurrency = schemaorg.NewProperty("discountCurrency")

	// The date that payment is due.
	PaymentDueDate = schemaorg.NewProperty("paymentDueDate")

	// Code used to redeem a discount.
	DiscountCode = schemaorg.NewProperty("discountCode")
)
