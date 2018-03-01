package orderitem

import "github.com/dpb587/go-schemaorg"

var (
	// The current status of the order item.
	OrderItemStatus = schemaorg.NewProperty("orderItemStatus")

	// The number of the item ordered. If the property is not set, assume the
	// quantity is one.
	OrderQuantity = schemaorg.NewProperty("orderQuantity")

	// The delivery of the parcel related to this order or order item.
	OrderDelivery = schemaorg.NewProperty("orderDelivery")

	// The item ordered.
	OrderedItem = schemaorg.NewProperty("orderedItem")

	// The identifier of the order item.
	OrderItemNumber = schemaorg.NewProperty("orderItemNumber")
)
