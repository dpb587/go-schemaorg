package itemlist

import "github.com/dpb587/go-schemaorg"

var (
	// For itemListElement values, you can use simple strings (e.g. "Peter", "Paul",
	// "Mary"), existing entities, or use ListItem.</p>
	// 
	// <p>Text values are best if the elements in the list are plain strings.
	// Existing entities are best for a simple, unordered list of existing things in
	// your data. ListItem is used with ordered lists when you want to provide
	// additional context about the element in that list or when the same item might
	// be in different places in different lists.</p>
	// 
	// <p>Note: The order of elements in your mark-up is not sufficient for
	// indicating the order or elements.  Use ListItem with a 'position' property in
	// such cases.
	ItemListElement = schemaorg.NewProperty("itemListElement")

	// Type of ordering (e.g. Ascending, Descending, Unordered).
	ItemListOrder = schemaorg.NewProperty("itemListOrder")

	// The number of items in an ItemList. Note that some descriptions might not
	// fully describe all items in a list (e.g., multi-page pagination); in such
	// cases, the numberOfItems would be for the entire list.
	NumberOfItems = schemaorg.NewProperty("numberOfItems")
)
