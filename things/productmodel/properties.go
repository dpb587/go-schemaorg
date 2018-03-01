package productmodel

import "github.com/dpb587/go-schemaorg"

var (
	// A pointer from a previous, often discontinued variant of the product to its
	// newer variant.
	PredecessorOf = schemaorg.NewProperty("predecessorOf")

	// A pointer from a newer variant of a product  to its previous, often
	// discontinued predecessor.
	SuccessorOf = schemaorg.NewProperty("successorOf")

	// A pointer to a base product from which this product is a variant. It is safe
	// to infer that the variant inherits all product features from the base model,
	// unless defined locally. This is not transitive.
	IsVariantOf = schemaorg.NewProperty("isVariantOf")
)
